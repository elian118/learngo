package something

import "fmt"

func GetNm() {
	const val string = "val"
	// 축약형 변수 선언 => 오직 func 안에서만 작동
	name := "elian" // var name string = "elian" -> 입력된 값의 타입으로부터 타입유추 -> string 간주
	name = "lynn"
	fmt.Println(val)  // val
	fmt.Println(name) // lynn
}
