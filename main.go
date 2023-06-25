package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("temp-images/")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Size())
	}
}
