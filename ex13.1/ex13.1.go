// 구조체 : 여러 필드를 묶어서 사용하는 타입
/*
예시
type Student struct {
	Name String
	Class int
	No int
	Score float64
}

이런 식으로 Student라는 구조체를 만들어주고 나면
var a Student 이런 식으로 a라는 학생의 속성을 정할 수 있다.
*/
package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구 ..."
	house.Size = 32
	house.Price = 10
	house.Category = "아파트"

	fmt.Println(house)
	// 위에꺼는 다음과 같은 것  fmt.Printf("%v\n", house)
	// 또 다음과더 같은 것     fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
}
