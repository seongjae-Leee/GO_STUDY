// 객체 : 데이터와 기능을 묶은 것으로 결합도를 높여줌. 메서드는 이 객체 사이의 관계를 정의함

// 포인터 타입 메서드와 값타입 메서드를 알아보자
package main

import "fmt"

type account struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account) WithdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) WithdrawValue(amount int) {
	a2.balance -= amount
}

// 만약 값을 깎고 싶다면 아래 함수를 써야 한다.
func (a3 account) WithdrawValue2(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{500, "seongjae", "Lee"}
	mainA.WithdrawPointer(30)
	fmt.Println(mainA.balance)
	// amount 50의 값이 WithdrawValue의 amount. 즉 a2가 깎이게 되고 mainA.balance에는 영향을 안 끼친다. 즉 값이 복사된 다른 mainA에서 깎는 것
	mainA.WithdrawValue(50)
	fmt.Println(mainA.balance)
	// 값이 깎이는 함수
	*mainA = mainA.WithdrawValue2(50)
	fmt.Println(mainA.balance)
}

// 결과값 : 470 470
// 결과값 : 470 470 420

// 그럼 언제 값타임을 쓰고 언제 포인터 타입을 쓰는지 다음장에서 알아보자.
