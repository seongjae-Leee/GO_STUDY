package main

import "fmt"

func main() {
	str := "Hello 성재"
	arr := []rune(str) // rune은 int32의 별칭
	// len은 문자 길이가 아니라 바이트 길이를 반환
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c\n", arr[i], arr[i], arr[i])
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
