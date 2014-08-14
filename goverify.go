package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jensraaby/goverify/hashing"
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
		fmt.Fprintf(os.Stderr, "Usage: %s (hash|check) PATH \n", os.Args[0])
		os.Exit(-1)
	}

	mode := args[0]
	fmt.Println("GoVerify: performing operation:", mode)

	path, err := checkPath(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid path. Must be absolute or relative path to a directory.\n", path)
		fmt.Fprintf(os.Stderr, "Detail: %s  \n", err)
	}

	switch mode {
	case "hash":
		fmt.Println("Let's hash!")
		hashing.HashDir(path)
	case "check":
		fmt.Println("Not ready!")
	default:
		fmt.Fprintf(os.Stderr, "Invalid operation: %s \n", mode)
	}
}

func checkPath(path string) (safePath string, err error) {
	safePath = filepath.Clean(path)
	s, err := os.Stat(path)
	// Print error with path information
	if e, ok := err.(*os.PathError); ok {
		fmt.Println("Error:", e)
		return
	}
	if os.IsNotExist(err) {
		//TODO Generate error for nonexisting
		// This is a type assertion. It allows us to assert that a variable is
		// not nil and the value stored within it is of a certain type. This is
		// how you do dynamic typing/checking of types
		err = err.(*os.PathError)
		//err = &GVError{"Check path", time.Now(), "Path does not exist"}
		return
	}
	if !s.IsDir() {
		// VerifyError{"checkPath", path, err}
		// e := fmt.Errorf("Not a directory: %s", path)
		// myerror := (&os.PathError{"Read", path, errors.New("Not a directory")})
		//err = myerror.(*os.PathError)
		// err = e.(*os.PathError)
		e := &GVError{"Check path", time.Now(), "Not a directory"}
		err = e
		return
	}
	return
}
