package something

import "fmt"

func CanIDrink(age int) bool {
	defer fmt.Println("function canIDrink done.")
	// 조건문에서 koreanAge 변수 선언과 동시에 명령 수행 -> 이 경우, koreanAge 변수는 조건문 안에서만 쓰임
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func CanIDrink2(age int) bool {
	defer fmt.Println("function canIDrink2 done.")
	koreanAge := age + 2
	return koreanAge >= 18
}

func CanIDrink3(age int) bool {
	defer fmt.Println("function canIDrink3 done.")
	// switch 역시 변수선언과 동시에 명령 수행 가능
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func CanIDrink4(age int) bool {
	defer fmt.Println("function canIDrink4 done.")
	// switch 문에 변수를 생략하면, case 문에 조건식 입력도 가능
	switch {
	case age+2 < 18:
		return false
	case age+2 == 18:
		return true
	case age+2 > 50:
		return false
	}
	return false
}

func ShowResultCaIDrink() {
	fmt.Println(CanIDrink(16))  // true
	fmt.Println(CanIDrink(15))  // false
	fmt.Println(CanIDrink2(16)) // true
	fmt.Println(CanIDrink2(15)) // false
	fmt.Println(CanIDrink3(16)) // true
	fmt.Println(CanIDrink3(15)) // false
	fmt.Println(CanIDrink4(16)) // true
	fmt.Println(CanIDrink4(15)) // false
}
