package utils

import (
	"fmt"
	"os"
	"sort"

	"github.com/dustin/go-humanize"
)

var totalSize float32

func PrintItemsByValue(dictionary map[string]int) {

	keys := make([]string, 0, len(dictionary))
	for key := range dictionary {
		keys = append(keys, key)
		totalSize += float32(dictionary[key])
	}

	file, fileErr := os.Create("file.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	fmt.Fprintf(file, "Total size:%vKB\n", humanize.Commaf(float64(totalSize)))

	sort.Slice(keys, func(i, j int) bool { return dictionary[keys[i]] > dictionary[keys[j]] })

	for _, key := range keys {

		fmt.Fprintf(file, "------------------------------------------------\n")

		fmt.Fprintf(file, "Filename: %s\n", key)
		fmt.Fprintf(file, "Size: %s KB | %% %.f of total space\n", humanize.Commaf(float64(dictionary[key])), float32(dictionary[key])/totalSize*100)

	}

}
