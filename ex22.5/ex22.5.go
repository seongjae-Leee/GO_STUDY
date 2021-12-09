// Ring : 뱀이 자기 꼬리를 물고 있듯 마지막에서 다시 처음으로 돌아오는 구조

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // ❶ 요소가 5개인 링 생성

	n := r.Len() // ❷ 링 개수 반환

	for i := 0; i < n; i++ {
		r.Value = 'A' + i // ❸ 순회하면 모든 요소에 값 대입
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value) // ❹ 순회하며 값 출력
		r = r.Next()
	}

	fmt.Println() // 한줄 띄우기

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value) // ➎ 순회하며 값 출력
		r = r.Prev()
	}
}
