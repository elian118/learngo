package something

import "fmt"

// private 함수는 소문자, public 함수는 대문자로 시작
func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}
