package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func stringToUpper(str string) string {
	var a string
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			a += string('A' + (v - 'a')) // Builder 이용하지 않고 합 연산하기 (대문자로 바꾸기)
		} else {
			a += string(v)
		}
	}
	return a
}

// strings.Builder 객체를 이용해서 문자를 더해주는 것
// strings.Builder는 내부의 slice를 가지고 있기 때문에 WriteRune을 통해 문자를 더할 때 매번 메모리를 생성하지 않고
// 기존 메모리 공간에 빈자리가 있으면 그냥 더한다. 따라서 공간낭비를 줄일 수 있다.
func stringToUpper1(str string) string {
	var b strings.Builder
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			b.WriteRune('A' + (v - 'a')) // strings.Builder 사용
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}

func main() {
	// 문자열 합산
	var str string = "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // 내부 구조체로 변경
	address1 := stringHeader.Data                                 // Data 필드 값을 변수로 저장
	str += "Seongjae"

	address2 := stringHeader.Data // Data 필드값을 변수로 저장
	str += "Hi"
	address3 := stringHeader.Data // Data 필드값을 변수로 저장
	fmt.Println(str)
	fmt.Printf("주소1 : \t%x\n", address1)
	fmt.Printf("주소2 : \t%x\n", address2)
	fmt.Printf("주소3 : \t%x\n", address3)
	/* ▲ 즉 기존 문자열 메모리 공간을 건드리지 않고 새로운 메모리 공간을 잡는다.
	이로써 문자열 불변의 원칙을 지킴을 알 수 있다. (메모리 낭비에도 버그가 나타나지 않게 하기 위해 불변)
	문자열 합 연산이 빈번하게 발생되면 메모리가 낭비됨. 만약 빈번하게 사용된다고 하면 string 패키지의 ★Builder를 이용해 메모리 낭비를 줄이도록 한다.
	*/
	var upperString string = "Hello World"
	fmt.Println(stringToUpper(upperString))
	fmt.Println(stringToUpper1(upperString))
}

/*
결과값 확인 방법 F5 눌러서 디버그 콘솔로 확인
결과값 :
HelloSeongjaeHi
주소1 : 	7dd602
주소2 : 	c000014090
주소3 : 	c0000140a0

HELLO WORLD
HELLO WORLD
*/
