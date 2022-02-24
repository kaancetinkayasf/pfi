package utils

import (
	"fmt"
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

	fmt.Printf("Total Size: %fKB\n", totalSize)
	sort.Slice(keys, func(i, j int) bool { return dictionary[keys[i]] > dictionary[keys[j]] })

	for _, key := range keys {

		fmt.Println("------------------------------------------------")

		fmt.Printf("Filename: %s\n", key)
		fmt.Printf("Size: %s KB | %% %.f of total space\n", humanize.Commaf(float64(dictionary[key])), float32(dictionary[key])/totalSize*100)

	}
}
