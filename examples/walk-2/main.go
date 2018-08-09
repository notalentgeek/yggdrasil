package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir       := "/home/notalentgeek/go"
	dirToSkip := ".git"

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// Prevent panic
		if err != nil {
			return err;
		}

		// Exclude directory
		if info.IsDir() && info.Name() == dirToSkip {
			fmt.Printf("Skipping Directory: %q\n", path)
			
			return filepath.SkipDir
		}

		fmt.Printf("Visited File      : %q\n", path)

		return nil
	})

	if err != nil {
		fmt.Printf("Error When Walking Down The Path %q: %v", dir, err)
	}
}