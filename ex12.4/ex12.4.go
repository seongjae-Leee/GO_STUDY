package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.7, 26.8, 26.2, 27.1}

	// range : 각 요소를 순회하며 범위 지정.
	// 즉 아래에서는 i랑 v를 반복하게 되는데 i에는 index 번호가 오고 v자리에는 t배열의 0번 요소부터 나옴.
	for i, v := range t {
		fmt.Println(i, v)
	}
}
