package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product) // ❶ 맵 생성

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	// 키, 값   range로 순회
	for k, v := range m { // ❸ 맵 순회
		fmt.Println(k, v)
	}
}

/*
결과값 :
16 {볼펜 500}
46 {지우개 200}
78 {자 1000}
345 {샤프 3000}
897 {샤프심 500}
-> 여기서는 잘 나왔지만 golang의 map은 순서 보장이 안됨
*/
