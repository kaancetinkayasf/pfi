package main

import (
	"fmt"
	"kaancetinkaya/m/pkg/utils"
	"os"
)

func main() {

	dictionary := make(map[string]int)
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	utils.ReadDir(path, dictionary)
	utils.PrintItemsByValue(dictionary)
}
