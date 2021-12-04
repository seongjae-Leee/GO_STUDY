package main

import (
	"fmt"
	"reflect" // 타입 내부를 살펴볼 때 필요
	"unsafe"
)

func main() {
	str1 := "Hello 성재"
	str2 := str1
	// str1의 메모리 주소를 unsafe의 포인터 타입으로 변환하는 것. 그것을 다시 reflect 패키지 안의 StringHeader 타입으로 변환해주는 것
	StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(StringHeader1)
	fmt.Println(StringHeader2)
}

/*
결과값 :
&{4613634 12} // 앞에꺼는 data 뒤에꺼는 len
&{4613634 12}
*/
