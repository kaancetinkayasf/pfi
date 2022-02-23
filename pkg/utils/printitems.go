package utils

import (
	"fmt"
	"sort"

	"github.com/dustin/go-humanize"
)

func PrintItemsByValue(dictionary map[string]int) {

	keys := make([]string, 0, len(dictionary))
	for key := range dictionary {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return dictionary[keys[i]] > dictionary[keys[j]] })

	for _, key := range keys {
		fmt.Printf("Filename: %s\n", key)
		fmt.Printf("Size: %s KB\n", humanize.Commaf(float64(dictionary[key])))
		fmt.Println("------------------------------------------------")
	}
}
