package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ignoredir string = ".git"
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
			// fmt.Println("Hash:", hashFile(path))
			fmt.Println("JPR: Passing to walker")
			printStuff(path)

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

/*
 Basically, exploring a directory structure is a matter of walking through a tree with a breadth-first-search. Go has a facility for this: the filepath.Walk function and WalkFunc type.
 Here I will try and use the latter to print the mod-time of everything in a dir.
*/
func mtimePrinter(path string, info os.FileInfo, err error) error {
	// there is a special error SkipDir we can use to avoid expanding dirs

	if bytes.HasPrefix([]byte(path), []byte(ignoredir)) {
		fmt.Println("WARNING: Path skipping, has ignoredir as prefix")
		fmt.Println("Ignoredir:", ignoredir)
		return filepath.SkipDir
	}
	f, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error stat'ing path,", path)
	}

	fmt.Println("Path", path, "Modification time:,", f.ModTime())
	return nil
}

func printStuff(startpath string) error {
	filepath.Walk(startpath, mtimePrinter)
	return nil
}
