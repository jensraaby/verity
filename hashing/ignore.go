// Package hashing provides the ability to skip folders based on a gitignore
// file.
//TODO Change gitignore to a custom file at somepoint.
package hashing

import "fmt"

var (
	Patterns []string
)

func LoadIgnoreFile(path string) error {
	// try to load the file
	// go through the lines and add to the array of patterns
	return nil
}

// Accepts: test whether a path is allowed according to our patterns
func Accepts(path string) bool {
	//TODO - fill in
	for pat := range Patterns {
		//do a check of the file, if anything is false then return
		fmt.Println("pat", pat)
	}
	return true
}
