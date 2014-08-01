package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"os"
	"strings"
	"time"
)

type FileHash struct {
	md5sum hash.Hash
	mtime  time.Time
}

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

	hasher := md5.New()
	var asbytes []byte
	asbytes = strings.Index(file)
	fmt.Println(hash.Sum(file))

	// use OS to get modification time
	f, err := os.Stat(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting file details for file: %s \n", file)
		os.Exit(-1)
	}
	h.mtime = f.ModTime()

	// do the hashing
	h.md5sum = hash.Sum(asbytes)
	return h
}
