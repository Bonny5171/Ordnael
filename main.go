package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func sniff(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
		sniff(file)
	}
}

func main() {
	log.Println(`
.
.
Ola! eu sou Ordnael
.
.
.
¯\_( ͡° ͜ʖ ͡°)_/¯`)

	sniff(".")
}
