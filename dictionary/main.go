package main

import (
	"dictionary/dictionary"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

	flag.Parse()

	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action %v\n", *action)
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("'%v' added to dictionary", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(strings.ToLower(word))
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(strings.ToLower(word))
	handleError(err)
	fmt.Printf("'%v' removed from dictionary", word)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
