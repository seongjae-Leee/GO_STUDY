// 쇼트 서킷으로 인해 생길 수 있는 문제를 파악해보자
package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	// ▼ 이 아래부분은 좌변이 false라서 우변을 검사 안하기 때문에 우변이 참이어도 "1 증가"가 출력되지 않는다.
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
}

// true이면
// IncreaseAndReturn() 0
// 1 증가                 라는 값이 나온다.

/*
쇼트 서킷을 활용하려면.. 예를 들어 식당 예약한 사람을 분류하고 싶을 때
if hasBooked() && 예약인원 > 3 {

}
이런 식으로 조건문을 적어두면 예약한 사람이 아닌 경우에는 좌변이 false이기 때문에 조건문 계산을 건너뛸 수 있는 것이다.
*/
