package main

import (
	"dictionary/dictionary"
	"fmt"
	"os"
)

func main() {
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

	_ = d.Add("golang", "a wonderful langage")
	_ = d.Add("python", "an interpreted langage")
	_ = d.Remove("python")
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
