package main

import (
	"fmt"
	// unsafe는 sizeof같은 타입의 사이즈 정보를 가져올 때 사용하는 패키지
	"unsafe"
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

// 포인터 타입으로 되어있는 String 메서드
func (s *Student) String() string {
	return s.Name
}

type User struct {
	Name string
	Age  int
}

// 값 타입으로 되어있는 String 메서드
func (u User) String() string {
	return u.Name
}

func main() {
	var stringer1 Stringer
	fmt.Printf("type: %T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	student := &Student{"AAA"}
	stringer1 = student
	fmt.Printf("type:%T size:%d\n", stringer1, unsafe.Sizeof(stringer1))

	var stringer2 Stringer
	fmt.Printf("type:%T size:%d\n", stringer2, unsafe.Sizeof(stringer2))

	user := User{"BBB", 20}
	stringer2 = user
}

/*
결과값 :
type: <nil> size:16
type:*main.Student size:16
type:<nil> size:16
*/
