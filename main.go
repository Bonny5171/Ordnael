package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func varrediretorio(root string) {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	log.Println(`
	.
	Ola! eu sou Ordnael
	.
	¯\_( ͡° ͜ʖ ͡°)_/¯`)

	varrediretorio(`C:\TesteOrdinael`)
}
