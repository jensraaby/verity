package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

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
		err = err.(*os.PathError)
		return
	}
	if !s.IsDir() {
		fmt.Println("Argh!")
		// VerifyError{"checkPath", path, err}
		e := fmt.Errorf("Not a directory: %s", path)
		// myerror := (&os.PathError{"Read", path, errors.New("Not a directory")})
		//err = myerror.(*os.PathError)
		// err = e.(*os.PathError)
		err = e
		return
	}
	fmt.Println(s.Name(), s.ModTime(), s.IsDir())
	return
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
