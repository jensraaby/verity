package hashing

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
)

// This is an implementation of concurrent hashing of files
// We assume that filepaths are received from a dir processer, which stores
// a slice of FileHashes (received on dest channel)

type Result struct {
	f   FileHash
	err error
}

func ProcessAllParallel(root string) ([]FileHash, error) {

	done := make(chan struct{})
	defer close(done)

	fhashes, errc := ProcessDir(done, root)

	m := make([]FileHash, 0)
	for fh := range fhashes {
		if fh.err != nil {
			return nil, fh.err
		}
		m = append(m, fh.f)
	}
	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}

func ProcessDir(done <-chan struct{}, dirpath string) (<-chan Result, <-chan error) {
	receive := make(chan Result)
	errc := make(chan error, 1)
	go func() {
		// waitgroup makes sure we don't close channels until all workers are
		// done
		var wg sync.WaitGroup
		err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				// this means its a file with no mode bits set
				// can also check IsDir with this
				return nil
			}
			wg.Add(1)
			// fire up a routine to read and handle the file
			go func() {
				select {
				case receive <- ProcessFile(path):
				case <-done:
				}
				wg.Done()
			}()
			// abort if done is closed
			select {
			case <-done:
				return errors.New("walk cancelled")
			default:
				return nil

			}

		})
		// walk has returned
		go func() {
			wg.Wait()
			close(receive)
		}()
		errc <- err
	}()
	return receive, errc
}

func ProcessFile(fpath string) Result {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return Result{FileHash{}, err}
	}
	stat, err := os.Stat(fpath)
	if err != nil {
		return Result{FileHash{}, err}
	}

	sum := sha1.Sum(data)
	mtime := stat.ModTime()
	sumstring := fmt.Sprintf("%x", sum)

	f := FileHash{path.Base(fpath), len(data), mtime, sumstring}
	return Result{f, nil}
}

// Work consumes filepaths on the source channel, and sends hash results to the
// dest channel
func Work(done <-chan struct{}, source <-chan string, dest chan FileHash) {
	for {
		select {
		case file := <-source:
			fmt.Println(file)
		case <-done:
			// done receiving on this worker
			return
		}
	}
}
