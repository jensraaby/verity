package hashing

import (
	"fmt"
	"os"
	"path/filepath"
)

func VerifyHashes(startpath string) (err error) {
	// first we use a walker function to build our file list, then we use
	// goroutines to check
	//type WalkFunc func(path string, info os.FileInfo, err error) error
	err = filepath.Walk(startpath, func(spath string, info os.FileInfo, incomingErr error) error {
		// if incoming err detected, let's see what it is:
		if incomingErr != nil {
			fmt.Fprintf(os.Stderr, incomingErr.Error())
		}
		s, localerr := os.Stat(spath)
		fmt.Println(spath, s.ModTime())
		return localerr
	})
	return
}

func CheckDir(dir string) (status bool, err error) {
	status = false
	// check if existing integrity data is present
	// if not, run hash and return true
	// otherwise, load old integrity file and scan through
	return
}
