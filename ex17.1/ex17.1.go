package main

import "fmt"

type account struct {
	balance int
}

// 1. Function 타입
// 첫번째 인자로는 account의 포인터 타입을, 두번째 인자로는 amount를 받는다.
func withdrawFunc(a *account, amount int) {
	a.balance -= amount // account 포인터의 balacne 필드값을 amount만큼 빼라.
}

// 2. 리시버 타입
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} // a의 타입은 포인터 타입

	withdrawFunc(a, 30)
	a.withdrawMethod(30) // a라는 객체(인스턴스)를 리시버로 받는 WithdrawMethod 메서드를 호출. 거기의 amount가 30

	fmt.Printf("%d\n", a.balance)
}

// 결과값 : 40
