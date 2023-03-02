package something

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

func RepeatMe(words ...string) {
	defer fmt.Println("function repeatMe done.")
	fmt.Println(words)
}

func ShowLenAndUpperSample() {
	totalLength, upperName := lenAndUpper("elianus")
	// totalLength, _ := lenAndUpper("elianus") // 언더스코어 -> 무시
	fmt.Println(totalLength, upperName) // 7 ELIANUS
	// fmt.Println(totalLength)                            // 7
	totalLength2, upperName2 := lenAndUpper2("elianus2") // 언더스코어 -> 무시
	fmt.Println(totalLength2, upperName2)                // 8 ELIANUS2
}
