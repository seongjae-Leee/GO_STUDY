package main

import "fmt"

func main() {
	str := "Hello 성재"
	// len은 문자 길이가 아니라 바이트 길이를 반환
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c\n", str[i], str[i], str[i])
	}
}

/*
결과값 :
타입 : uint8 값 : 72 문자값 : H
타입 : uint8 값 : 101 문자값 : e
타입 : uint8 값 : 108 문자값 : l
타입 : uint8 값 : 108 문자값 : l
타입 : uint8 값 : 111 문자값 : o
타입 : uint8 값 : 32 문자값 :
타입 : uint8 값 : 236 문자값 : ì
타입 : uint8 값 : 132 문자값 :

타입 : uint8 값 : 177 문자값 : ±
타입 : uint8 값 : 236 문자값 : ì
타입 : uint8 값 : 158 문자값 : 입 : uint8 값 : 172 문자값 : ¬
*/
