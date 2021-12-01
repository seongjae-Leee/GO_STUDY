// 이중배열 (다중배열로 가면 tree 구조처럼.. 뭔지 알지?)
package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 이 아래 10번째 줄 중괄호가 줄바꿈이 되어있기 때문에 콤마 필요
	}

	for _, arr := range a { // 0, 1 이 추출됨
		for _, v := range arr {
			fmt.Print(v, ",") // a[0]과 a[1] 값이 추출됨
		}
		fmt.Println()
	}
}
