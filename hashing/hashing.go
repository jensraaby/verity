// Package hasher provides functionality to gather file hashes for all files in
// a directory. It will either save to a specified file or write one to each
// directory
package hashing

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// This file handles the hashing operations.
// 1. Start with a path as input
// 2. Pass to a walker which sends all files to a channel
// 3. Channel receiver puts files in a datastructure
// 4. Datastructure is passed to a queue to process the hashes

type FileHash struct {
	name    string
	modtime time.Time
	sha1    []byte
}
type DirHash struct {
	path  string
	files []FileHash
}

func HashDir(path string) error {
	fmt.Println(path)
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
	return nil
}
