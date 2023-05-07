package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("I can only find one item at a time")
		os.Exit(1)
	}

	if err := run(args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(fileName string) error {
	return filepath.WalkDir(".", func (path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == fileName {
			if d.IsDir() {
				fmt.Println("found: " + path + d.Name())
			} else {
				fmt.Println("found: " + d.Name())
			}
		}
		
		return nil
	})
}