// ● 메서드 : 함수의 일종.. 객체랑 묶어주는 녀석. 즉 데이터와 기능을 묶어주는 단일 객체와 묶어줄 수 있게 됨.
// go에서는 C++과 다르게 class가 없기 때문에 구조체 밖에 메서드를 지정해줘야 함.
// 리시버(receiver)를 통해 메서드가 어느 구조체에 속하는지 소속을 나타냄.
package main

import "fmt"

// 타입으로 선언된 것들을 리시버로 사용함.
type account struct {
	b int
}

// 일반함수 표현 방식
func Function(a *account, amount int) {
	a.b -= amount
}

// 이걸 메서드 표현 방식으로 하면... Method라는 메서드가 리시버인 *account에 속한다는 뜻
func (a *account) Method(amount int) {
	a.b -= amount
}

func main() {

	// ▼ 메서드 선언하기
	//      리시버  메서드이름 ... cal이라는 메서드가 Rect 타입에 속한다고 알려주는 것.
	// func(r Rect) cal() int{
	// 	return r.width * r.height
	// }

	//      account의 주소
	test := &account{200} //b가 200
	Function(test, 30)

	// Method가 account에 속하니까 test.Method가 가능한 것
	test.Method(20)

	fmt.Printf("%d\n", test.b)
}

// 결과값 : 200 - 30 - 20 = 150
