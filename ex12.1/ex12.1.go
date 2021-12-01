// 배열 선언시 요소의 개수는 항상 상수
package main

import "fmt"

func main() {
	// t라는 공간에 5개 element로 이루어진 float64 타입의 배열을 넣어주는 상황.
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}

/*
결과값 :
24
25.9
27.8
26.9
26.2
*/

/*
다양한 변수 선언
1. var nums [5]int  -> 5개 intiger로 구성된 배열
2. days := [3]string{"monday", "tuesday", "wednesday"}
3. var temps [5]float64 = [5]float64{24.3, 26.7}  -> 나머지 세개는 0.0으로 채워짐
4. var s =[5]int{1:10, 3:30}  -> {0, 10, 0, 30, 0}
5. x := [...]int{10, 20, 30}  -> 길이가 뒤에 있는 요소의 초기값(갯수)에 따라 배열의 길이가 정해짐. 즉 x 라는 배열은 앞으로도 element 3개
6. y := [...]int{10, 20, 30}  -> 길이가 계속 변할 수 있음
*/
