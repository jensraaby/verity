package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"time"
)

// Structure to represent the hash and last mod time of a file
type FileHash struct {
	md5sum [16]byte
	mtime  time.Time
}

// Load a file from a path into a byte slice
func LoadFile(file string) []byte {
	// we have already checked that the file exists, but handle it here too
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}

	s, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	data := make([]byte, s.Size())
	_, readErr := f.Read(data)
	if readErr != nil {
		fmt.Fprintf(os.Stderr, readErr.Error())
		os.Exit(-1)
	}
	return data
}

func getHash(file string) *FileHash {
	// load the file, call md5
	h := new(FileHash)

	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
	hasher := md5.New()
	io.Copy(hasher, f)

	h.md5sum = md5.Sum(nil)

	// use OS to get modification time
	st, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting file details for file: %s \n", file)
		os.Exit(-1)
	}
	h.mtime = st.ModTime()

	// do the hashing
	return h
}
