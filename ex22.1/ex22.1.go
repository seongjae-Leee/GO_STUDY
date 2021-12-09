// 자료구조 : 자료들을 어떤 형태로 저장할 것인가를 나타냄. 배열, 리스트, 트리, 맵 등이 있음
// 리스트 : 배열과 함께 가장 기본적인 선형 자료구조 중 하나
/* 기본구조
type Element struct{ // 구조체
	Value interface{} // 데이터를 저장하는 필드
	Next *Element 		// 다음 요소의 주소를 저장하는 필드
	Prev *Element 		// 이전 요소의 주소를 저장하는 필드(Double Linked List)
}
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4)   // 리스트 맨 뒤에 4 -> 4
	e1 := v.PushFront(1)  // 리스트 맨 앞에 1 -> 1 4
	v.InsertBefore(3, e4) // e4 앞에 3 -> 1 3 4
	v.InsertAfter(2, e1)  // e1 뒤에 2 -> 1 2 3 4

	//         ▼ 맨 앞 element 반환     ▼ 다음 노드로
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	//         ▼ 맨 뒤 element 반환     ▼ 그 전 노드로
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}

/*
결과값 :
1 2 3 4
4 3 2 1
*/
