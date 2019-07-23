package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Println(path)
	}
	return err
}

func main() {

	root := "./"

	err := filepath.Walk(root, walkFunc)

	fmt.Println(err)

}
