package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"github.com/elian/learngo/accounts"
	"github.com/elian/learngo/something"
)

func theory() {
	fmt.Println("--------variable-------")
	something.GetNm()
	fmt.Println("--------calculate-------")
	fmt.Println(something.Multiply(2, 2)) // 4
	fmt.Println("-------word--------")
	something.ShowLenAndUpperSample()
	fmt.Println("-------word--------")
	something.RepeatMe("nico", "lynn", "dal", "marl", "flynn") // [nico lynn dal marl flynn]
	fmt.Println("--------calculate-------")
	something.ShowSuperAddSample()
	fmt.Println("--------condition-------")
	something.ShowResultCaIDrink()
	fmt.Println("--------variable-------")
	something.ShowGetMemorizedInfo()
	fmt.Println("--------array-------")
	something.ShowArr()
	fmt.Println("--------map-------")
	something.ShowMap()
	fmt.Println("--------structs == object-------")
	something.ShowStruct()
}

func main() {
	// theory()
	account := accounts.NewAccount("elian")
	fmt.Println(account) // &{elian 0}
	// 물론, Account 각 필드는 모두 private 속성(소문자)이라, 아래 코드 실행 불가
	// accounts.balance = 100000
	// accounts.owner = "pepe"
	account.Deposit(100000000)
	fmt.Println(account.Balance()) // 0 -> 변경되지 않았다.
}
