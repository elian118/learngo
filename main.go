package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"github.com/elian/learngo/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello()
}
