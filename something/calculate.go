package something

import "fmt"

func Multiply(a int, b int) int {
	defer fmt.Println("function multiply done.")
	return a * b
}

func superAddEx(numbers ...int) int {
	defer fmt.Println("function superAddEx done.")
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
	defer fmt.Println("function superAdd done.")
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func ShowSuperAddSample() {
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
}
