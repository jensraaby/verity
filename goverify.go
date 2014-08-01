package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// todo: create arguments here using flags package
//sourceDir := *flag.Flag{name,usage,value,defvalue}
func main() {
	// Check incoming arguments. We need 2 paths
	//var source = flag.String("source", ".", "source directory")
	//var dest = flag.String("dest", ".", "dest directory")
	flag.Parse()
	args := flag.Args()

	//fmt.Println("source arg:", *source)
	//fmt.Println("dest arg:", *dest)

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ORIGIN DESTINATION \n", os.Args[0])
		os.Exit(-1)
	}

	// run through the args..
	for _, path := range args {
		fmt.Println("Testing " + path)
		if exists(path) {
			fmt.Println("Is it absolute? ")
			fmt.Println(filepath.IsAbs(path))
			fmt.Println("Hash:", hashFile(path))
		}
		//data := filehash.LoadFile(path)
	}
}
func hashFile(path string) *FileHash {
	//h := &FileHash{}
	h := getHash(path)
	return h
}

//TODO: check there is no better way to do this
func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", path)
		return false
	}
	return true
}
