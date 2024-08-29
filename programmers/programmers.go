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

func Sol181941(arr []string) string {
	var result string

	for _, data := range arr {
		result += data
	}
	return result
}

// Sol181924 arr 에 들어있는 배열을 queries 를 통해 변경하여 반환
func Sol181924(arr []int, queries [][]int) []int {
	//  길이가 같은 새로운 배열생성
	result := make([]int, len(arr))
	// 배열 복사 (new, origin)
	copy(result, arr)
	for _, query := range queries {
		first := query[0]
		second := query[1]
		result[first], result[second] = result[second], result[first]
	}
	return result
}
