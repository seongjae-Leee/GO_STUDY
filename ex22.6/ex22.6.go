// 맵은 키와 값 형태로 데이터를 저장하는 자료구조이다. 언어에 따라서 딕셔너리, 해쉬테이블, 해쉬맵 등으로 부른다.
// 기본형 : map[key]value -> 예를 들어 map[string]int 라고 하면 key는 문자값이고 value는 정수값이라는 뜻.

package main

import "fmt"

func main() {
	m := make(map[string]string) // ❶ 맵 생성
	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	m["최번개"] = "청주시 상당구" // ❸ 값 변경

	fmt.Printf("송하나의 주소는 %s입니다.\n", m["송하나"]) // 송하나 라는 키값에 해당하는 value를 보여줄 것
	fmt.Printf("백두산의 주소는 %s입니다.\n", m["백두산"])
}
