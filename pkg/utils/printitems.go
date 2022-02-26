package utils

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	suffixes [4]string
)

var totalSize float64

func PrintItemsByValue(dictionary map[string]float64) {

	keys := make([]string, 0, len(dictionary))
	for key := range dictionary {
		keys = append(keys, key)
		totalSize += float64(dictionary[key])
	}

	file, fileErr := os.Create("results.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	fmt.Fprintf(file, "Total size: %s\n", HumanFileSize(totalSize))

	sort.Slice(keys, func(i, j int) bool { return dictionary[keys[i]] > dictionary[keys[j]] })

	for _, key := range keys {
		size := HumanFileSize(dictionary[key])
		fmt.Fprintf(file, "------------------------------------------------\n")

		fmt.Fprintf(file, "Filename: %s\n", key)
		fmt.Fprintf(file, "Size: %s | %% %.2f of total space\n", size, (dictionary[key]/totalSize)*100)

	}

	fmt.Println("The results.txt file is created")
}

func HumanFileSize(size float64) string {

	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)

	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
