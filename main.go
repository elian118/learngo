package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"github.com/elian/learngo/accounts"
	"github.com/elian/learngo/common"
	"github.com/elian/learngo/mydict"
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

func bank() {
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

func dict() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "Hello"
	def := "Greeting"
	baseWord := "hello"

	fmt.Println("----- 1) add to use map")
	fmt.Println(dictionary) // map[first:First word]
	dictionary["second"] = "hello"
	fmt.Println(dictionary) // map[first:First word second:hello]
	fmt.Println("----- 2) add and search to use method")
	definition, err := dictionary.Search("third")
	common.ErrHandler2(definition, err) // Not Found
	err = dictionary.Add(word, def)
	fmt.Println(dictionary)
	common.ErrHandler(err) // 에러가 없어 건너 뜀
	err = dictionary.Add(word, def)
	common.ErrHandlerForDict(word, definition, err) // That word already exists
	definition, err = dictionary.Search(word)
	common.ErrHandlerForDict(word, definition, err) // found: Hello / definition: Greeting
	fmt.Println("----- 3) update and search to use method")
	_ = dictionary.Add(baseWord, "Forth")
	fmt.Println(dictionary)
	err = dictionary.Update(baseWord, "Forth - updated")
	fmt.Println(dictionary)
	common.ErrHandler(err)
	value, _ := dictionary.Search(baseWord)
	fmt.Println(value)
	err = dictionary.Update("Fifth", "Fifth - updated 2") // Can't update non-existing word
	common.ErrHandler(err)
	fmt.Println("----- 4) delete and search to use method")
	err = dictionary.Delete("Fifth")
	common.ErrHandler(err) // Can't delete non-existing word
	_ = dictionary.Delete(baseWord)
	fmt.Println(dictionary)
	definition, err = dictionary.Search(baseWord)
	common.ErrHandlerForDict(baseWord, definition, err) // Not Found
}

func main() {
	dict()
}
