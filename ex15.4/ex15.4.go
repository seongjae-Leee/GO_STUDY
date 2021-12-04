package main

import "fmt"

func main() {
	str := "Hello 성재"
	// range라는 걸 for문 대신 쓸 수 있는 것.(index, 값) 에서 index 생략
	for _, v := range str {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c\n", v, v, v)
	}
}

/*
결과값 :
타입 : int32 값 : 72 문자값 : H
타입 : int32 값 : 101 문자값 : e
타입 : int32 값 : 108 문자값 : l
타입 : int32 값 : 108 문자값 : l
타입 : int32 값 : 111 문자값 : o
타입 : int32 값 : 32 문자값 :
타입 : int32 값 : 49457 문자값 : 성
타입 : int32 값 : 51116 문자값 : 재
*/
