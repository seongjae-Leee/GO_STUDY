package main

import "fmt"

// 타입앞에 ... 이 찍히면 그 타입에 해당하는 것을 여러개 만들 수 있다는 뜻. 슬라이스 타입!!
func sum(nums ...int) int {
	sum := 0
	fmt.Printf("nums 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Print(sum(1, 2, 3, 4, 5))
	fmt.Print(sum(10, 20))
	fmt.Print(sum())
}

/*
결과값 :
nums 타입 : []int
15nums 타입 : []int
30nums 타입 : []int
0
*/
// 만약 ...int가 아니라 int 였다면 아래 main 함수에 sum([]int{1, 2, 3, 4, 5}), sum([]int{10, 20}), sum([]int{}) 이렇게 적어야 한다.
