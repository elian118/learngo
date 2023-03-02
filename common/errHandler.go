package common

import "fmt"

func ErrHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ErrHandler2(definition string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}

func ErrHandlerForDict(word string, definition string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("found:", word, "/", "definition:", definition)
	}
}
