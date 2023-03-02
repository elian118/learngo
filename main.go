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

func bankAndDictionary() {
	// theory()
	account := accounts.NewAccount("elian")
	fmt.Println(account) // &{elian 0}
	// 물론, Account 각 필드는 모두 private 속성(소문자)이라, 아래 코드 실행 불가
	// accounts.balance = 100000
	// accounts.owner = "pepe"
	account.Deposit(10)
	fmt.Println(account.Balance()) // 10
	_ = account.Withdraw(5)
	fmt.Println(account.Balance()) // 5
	err := account.Withdraw(20)
	// 에러핸들러
	if err != nil {
		// log.Fatalln(err) // 2023/03/02 13:30:36 Can't withdraw. -> 에러 띄우고 프로그램 종료
		fmt.Println(err) // Can't withdraw. -> 프로그램 실행 유지(예외처리)

	}
	fmt.Println(account.Balance(), account.Owner()) // 5 elian
	account.ChangeOwner("nico")
	fmt.Println(account.Balance())                  // 5
	fmt.Println(account.Balance(), account.Owner()) // 5 nico
	fmt.Println(account)                            // &{nico 5} -> String 오버라이드 -> nico's account. Has: 5

}

func main() {
	bankAndDictionary()
}
