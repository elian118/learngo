package something

import "fmt"

func arrEx() {
	names := [5]string{"elian", "nico", "dal"}
	fmt.Println(names) // [elian nico dal  ]
	names[3] = "lala3"
	names[4] = "lala4"
	// names[5] = "lala5" // 배열길이를 초과해 에러 발생
	fmt.Println(names) // [elian nico dal lala3 lala4]
}

func simpleArrEx() {
	// Go 배열은 초기에 길이 제한이 없지만, 선언과 동시에 요소 수에 따라 고정 제한길이가 지정된다.
	names := []string{"elian", "nico", "dal"} // 길이가 3인 배열 선언
	fmt.Println(names)                        // [elian nico dal  ]
	// names[3] = "lala3" // 4번째 배열을 지정할 수 없어 에러 발생
	// 대신 아래처럼 새로운 배열을 선언해 기존 names 배열을 재정의하면 새로운 요소를 추가할 수 있다.
	// 	=> 즉, Go 배열은 자바스크립트의 Array.push('x') 처럼 원본 배열을 직접 변형하는 명령이 없다.
	names = append(names, "flynn")
	fmt.Println(names) // [elian nico dal flynn]
}

func ShowArr() {
	fmt.Println("1) arrEx")
	arrEx()
	fmt.Println("2) simpleArrEx")
	simpleArrEx()
}
