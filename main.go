package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	const val string = "val"
	// 축약형 변수 선언 => 오직 func 안에서만 작동
	name := "elian" // var name string = "elian" -> 입력된 값의 타입으로부터 타입유추 -> string 간주
	name = "lynn"
	fmt.Println("---------------")
	fmt.Println(val)  // val
	fmt.Println(name) // lynn
	fmt.Println("---------------")
	fmt.Println(multiply(2, 2)) // 4
	fmt.Println("---------------")
	//totalLength, upperName := lenAndUpper("elianus")
	totalLength, _ := lenAndUpper("elianus") // 언더스코어 -> 무시
	//fmt.Println(totalLength, upperName) // 7 ELIANUS
	fmt.Println(totalLength) // 7
	fmt.Println("---------------")
	repeatMe("nico", "lynn", "dal", "marl", "flynn") // [nico lynn dal marl flynn]
}
