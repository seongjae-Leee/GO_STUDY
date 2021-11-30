package main

import "fmt"

func main() {
	var age = 27

	if age >= 10 && age <= 15 {
		fmt.Println("You are young")
	} else if age > 30 || age < 20 {
		fmt.Println("You're not 20's")
	} else {
		fmt.Println("Best Age")
	}
}

// &&조건 연산자를 사용한 if문에서는 좌변에 false 확률이 높은? 조건을 걸어두는게 나을 것 같다
// 왜냐하면 어차피 좌변이 false 면 우변은 검사할 필요가 없이 false 이기 때문이고
// || 연산자를 사용한 if문에서는 좌변에 true 확률이 높은? 조건을 걸어두는게
// 왼쪽만 true 여도 우변을 검사하지 않고 true값을 내뱉기 때문에 효율적이다. (쇼트 서킷)
// 그러나 이로 인해 생길 수 있는 문제도 존재한다. 예제 9.4에서 확인해보자

// go 중급강의까지 이번주에 끝내고 nodejs랑 면접질문 공부 달리자
