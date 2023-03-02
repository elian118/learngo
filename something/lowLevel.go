package something

import "fmt"

func getRef() {
	a := 2
	b := a
	fmt.Println(a) // 2
	fmt.Println(b) // 2
	b = 10
	fmt.Println(a) // 2
	fmt.Println(b) // 10
}

// &변수: 변수가 저장된 메모리 주소를 가져온다.
// *변수: 변수가 메모리 주소라면, 그 메모리 주소를 참조해 그곳에 저장된 정보를 불러온다.
func getRef2() {
	// 다른 변수에 같은 메모리 주소를 참조시켜 불필요한 메모리 낭비를 줄이는 방식
	a := 2
	b := &a         // 이때부터 b는 메모리 주소정보를 담는 *int 타입이 된다. => int 타입으로 변경 불가
	fmt.Println(a)  // 2
	fmt.Println(&a) // 0x14000126018
	fmt.Println(b)  // 0x14000126018
	fmt.Println(*b) // 2
	a = 5
	fmt.Println(*b) // 5
	*b = 20         // 메모리에 저장된 값 변경
	fmt.Println(a)  // 20
}

func ShowGetMemorizedInfo() {
	fmt.Println("1) get memorized value")
	getRef()
	fmt.Println("2) get memory address")
	getRef2()
}
