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

// go 중급강의까지 이번주에 끝내고 nodejs랑 면접질문 공부 달리자
