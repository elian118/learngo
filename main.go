package main // 컴퓨터가 컴파일하는 파일은 오직 main 패키지 뿐이다.

import (
	"fmt"
	"github.com/elian/learngo/something"
)

func main() {
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
}
