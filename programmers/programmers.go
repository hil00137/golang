package programmers

import (
	"fmt"
	"unicode"
)

// Sol181952 대문자로 시작할경우 Public 함수
func Sol181952() {
	var s1 string
	// 입력된 값을 s1에 할당
	fmt.Scan(&s1)
	// 출력
	fmt.Println(s1)
}

func Sol181949() {
	var s1 string
	fmt.Scan(&s1)
	// String -> []rune 로 전환
	// rune 이란 유니코드 문자처리를 위한 데이터 타입 (Unicode Code Point)
	result := []rune(s1)
	for index, char := range result {
		if unicode.IsLower(char) {
			result[index] = unicode.ToUpper(char)
		} else {
			result[index] = unicode.ToLower(char)
		}
	}
	fmt.Println(string(result))
}
