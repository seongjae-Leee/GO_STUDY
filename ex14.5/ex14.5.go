package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u // u의 주소를 반환
}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}

/*
결과값 :
&{AAA 23}
*/

// Escape Analyzing : 컴파일러가 코드를 분석해서 어떤 인스턴스가 func 밖으로 탈출하는지 아닌지 분석.
// -> 탈출하는 놈들을 스택이 아니라 힙에다가 만든다. 그래서 쓰임을 다하면 사라진다. 즉 쓰임이 있는 한 사라지지 않는다.
