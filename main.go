package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("function lenAndUpper done.") // defer -> 함수 종료 후 의무 실행
	return len(name), strings.ToUpper(name)
}

// lenAndUpper()와 결과 동일
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("function lenAndUpper2 done.") // defer -> 함수 종료 후 의무 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAddEx(numbers ...int) int {
	// Go 에서 for 문은 range 를 통해 최소 두 개 인자를 받게 돼 있으며,
	// 첫 번째 인자가배열의 인덱스,
	// 두 번째 인자가 실제 요소값이다. *
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	// 물론, range 없이 고전적인 방법을 써도 상관 없다.
	// for i := 0; i < len(numbers); i++ {
	// 	 fmt.Println(numbers[i])
	// }
	return 1
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// 조건문에서 koreanAge 변수 선언과 동시에 명령 수행 -> 이 경우, koreanAge 변수는 조건문 안에서만 쓰임
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	koreanAge := age + 2
	return koreanAge >= 18
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
	totalLength, upperName := lenAndUpper("elianus")
	// totalLength, _ := lenAndUpper("elianus") // 언더스코어 -> 무시
	fmt.Println(totalLength, upperName) // 7 ELIANUS
	// fmt.Println(totalLength)                            // 7
	totalLength2, upperName2 := lenAndUpper2("elianus2") // 언더스코어 -> 무시
	fmt.Println(totalLength2, upperName2)                // 8 ELIANUS2
	fmt.Println("---------------")
	repeatMe("nico", "lynn", "dal", "marl", "flynn") // [nico lynn dal marl flynn]
	fmt.Println("---------------")
	superAddEx(1, 2, 3, 4, 5, 6)
	/*
		0 1
		1 2
		2 3
		3 4
		4 5
		5 6
	*/
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result) // 21
	fmt.Println("---------------")
	fmt.Println(canIDrink(16))  // true
	fmt.Println(canIDrink(15))  // false
	fmt.Println(canIDrink2(16)) // true
	fmt.Println(canIDrink2(15)) // false
}
