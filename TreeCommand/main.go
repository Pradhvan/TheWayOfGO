package main

import (
	"fmt";
	"os";
	"log";
	"path/filepath"
)

func tree(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	return err
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	fmt.Println(args)

	for _, arg := range args {
		err := tree(arg)
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}