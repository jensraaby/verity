// Package hasher provides functionality to gather file hashes for all files in
// a directory. It will either save to a specified file or write one to each
// directory
package hashing

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
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
	fmt.Printf("% x\n", sha1.Sum(data))
	return nil
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
	filepath.Walk(startpath, mtimePrinter)
	return nil
}
