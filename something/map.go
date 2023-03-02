package something

import "fmt"

func mapEx() {
	elian := map[string]string{"name": "elian", "age": "39"}
	for key, value := range elian {
		fmt.Println(key, value)
	}
	/*
		name elian
		age 39
	*/
}

func ShowMap() {
	mapEx()
}
