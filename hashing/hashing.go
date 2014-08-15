// Package hasher provides functionality to gather file hashes for all files in
// a directory. It will either save to a specified file or write one to each
// directory
package hashing

// This file handles the hashing operations.
// 1. Start with a path as input
// 2. Pass to a walker which sends all files to a channel
// 3. Channel receiver puts files in a datastructure
// 4. Datastructure is passed to a queue to process the hashes

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// This file handles the hashing operations.
// 1. Start with a path as input
// 2. Pass to a walker which sends all files to a channel
// 3. Channel receiver puts files in a datastructure
// 4. Datastructure is passed to a queue to process the hashes

// FileHash is the internal representation of a filehash
type FileHash struct {
	Name     string
	Modtime  time.Time
	CheckSum string
}

// A DirHash represents all the files in a dir (does not include subdirs)
type DirHash struct {
	Path      string
	LastCheck time.Time
	Files     []FileHash
}

// HashDir takes a path (assuming it is correct!) and begins the process of
// hashing the files within
func HashDir(path string) error {
	// this is just a dumb printer
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if info.IsDir() {
			fmt.Println("Is a dir!")
			if strings.HasPrefix(path, ".git") {
				fmt.Println("It's a git")
			}
		}
		return nil
	})
	// h, err := hashFile("hashing/hashing.go")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// This function will return the SHA1sum of a file
func hashFile(fpath string) (hash FileHash, err error) {
	// a hash.Hash implementer is a Writer (you can send it a stream of bytes)
	// It has a Sum method which takes the current stream and returns the
	// result of the hash
	hasher := sha1.New()

	// File IO
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	s, err := os.Stat(fpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	mtime := s.ModTime()
	// We close the file after copying it to the hasher
	defer f.Close()
	// io.Copy takes a Writer and Reader. We are putting the file in the shasum
	// buffer
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
		return hash, err
	}
	// I have no idea why you need to pass in nil here:
	checkSum := hasher.Sum(nil)

	hash = FileHash{
		Name:     path.Base(fpath),
		Modtime:  mtime,
		CheckSum: fmt.Sprintf("%x", checkSum),
	}
	return
}

/*
 Basically, exploring a directory structure is a matter of walking through a tree with a breadth-first-search. Go has a facility for this: the filepath.Walk function and WalkFunc type.
 Here I will try and use the latter to print the mod-time of everything in a dir.
*/
func mtimePrinter(path string, info os.FileInfo, err error) error {
	// there is a special error SkipDir we can use to avoid expanding dirs

	// if bytes.HasPrefix([]byte(path), []byte(ignoredir)) {
	// 	fmt.Println("WARNING: Path skipping, has ignoredir as prefix")
	// 	fmt.Println("Ignoredir:", ignoredir)
	// 	return filepath.SkipDir
	// }
	f, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error stat'ing path,", path)
	}

	fmt.Println("Path", path, "Modification time:", f.ModTime())
	return nil
}

func printStuff(startpath string) error {
	return filepath.Walk(startpath, mtimePrinter)
}
