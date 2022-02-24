package main

import (
	"flag"
	"fmt"
	"os"
	"pfi/pkg/utils"
)

func main() {

	var path string
	dictionary := make(map[string]int)

	flag.StringVar(&path, "path", ".", "project path")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a path. If you want to inspect the current directory, try 'pfi -path .'")
		os.Exit(1)
	}

	utils.ReadDir(path, dictionary)
	utils.PrintItemsByValue(dictionary)

}
