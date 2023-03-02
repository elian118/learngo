package something

import "fmt"

func structEx() {
	// 참고: Go 언어는 생성자라는 개념이 없고 그런 슈퍼메소드도 존재하지 않는다.
	type person struct {
		name    string
		age     int
		favFood []string
	}

	_favFood := []string{"kimchi", "ramen"}
	// elian := person{"elian", 18, favFood}
	elian := person{name: "elian", age: 18, favFood: _favFood}
	fmt.Println(elian.name) // elian
	fmt.Println(elian)      // {elian 18 [kimchi ramen]}
}

func ShowStruct() {
	structEx()
}
