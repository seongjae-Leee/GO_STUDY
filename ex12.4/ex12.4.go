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

/*
결과값 :
0 24
1 25.7
2 26.8
3 26.2
4 27.1

※ 근데 여기서 v만 출력할 순 없음.
v값만 출력해서 요소의 값만 가져오고 싶으면 for i, v ... 대신에 for _, v ... 형식으로 쓰면  v만 프린트 가능
*/
