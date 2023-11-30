package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	// "fs
	"os"
)

func main() {
	path, err := getDir()
	if err != nil {
		fmt.Printf("error occured: %v \n", err)
		return
	}

	fmt.Printf("directory: %v \n", path)

	getTree(path)
}

func getDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("error occured: %v \n", err)
		return "", err
	}

	dir := filepath.Dir(currentDir)
	return dir, nil
}

func getTree(path string) error {
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error occured: %v \n", err)
			return err
		}
		if info.IsDir() {
			fmt.Println("Dir: ", path)
		} else {
			fmt.Println("File: ", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error occured: %v \n", err)
	}
	return nil
}
