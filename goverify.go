package main

import (
	"flag"
	//  "path"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check incoming arguments. We need 2 paths
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ORIGIN DESTINATION \n", os.Args[0])
		os.Exit(-1)
	}

	for _, path := range args {
		fmt.Println("Testing " + path)
		if exists(path) {
			fmt.Println("Is it absolute? ")
			fmt.Println(filepath.IsAbs(path))
		}
	}
}

//TODO: check there is no better way to do this
func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", path)
		return false
	}
	return true
}
