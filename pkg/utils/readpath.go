package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadDir(dirname string, dic map[string]int) (dict map[string]int) {

	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range list {

		if item.Name() == ".git" {
			continue
		}

		if item.IsDir() {

			ReadDir(dirname+"/"+item.Name(), dic)
			continue
		}

		if item.Size() < 1000 {
			continue
		}

		dic[item.Name()] = int(item.Size())

	}

	return dic
}
