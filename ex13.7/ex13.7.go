package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1바이트
	B int  //8바이트
	C int8
	D int
	E int8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

// 결과값은 19가 아니라 40이 나오고 이것 역시 memory padding 때문이다.
// 그렇다면 이를 없애기 위해서는 순서를 바꿔줘야 한다.
// 즉 같은 memory를 가진 것끼리(A, C, E) 먼저 해주면 memory padding을 줄여서 24바이트가 나온다.
