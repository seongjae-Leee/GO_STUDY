package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 200
}
func changeSlice(slice []int) {
	slice[2] = 200
}

func main() {
	// type SliceHeader struct {
	// 	Data uintptr // 실제 배열을 가르키는 포인터
	// 	Len  int     // 요소 개수
	// 	Cap  int     // 실제 배열의 길이
	// } -> 슬라이스의 구조

	array := [5]int{1, 2, 3, 4, 5}
	sli := []int{1, 2, 3, 4, 5}

	changeArr(array)
	changeSlice(sli)
	fmt.Println(array)
	fmt.Println(sli)
	/*
		[1 2 3 4 5]
		[1 2 200 4 5]
	*/

	slice1 := make([]int, 3, 5) // len:3 cap:5의 슬라이스 만듦
	slice2 := append(slice1, 4, 5)
	fmt.Println("슬라이스1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("슬라이스2 : ", slice2, len(slice2), cap(slice2))
	/*
		슬라이스1 :  [0 0 0] 3 5
		슬라이스2 :  [0 0 0 4 5] 5 5

		만약 빈공간이 없으면 더 큰 배열을 만든다. 일반적으로 기존 배열의 2배짜리로.
		-> 기존 배열의 요소를 새로운 배열에 복사 -> 새로운 배열 맨 뒤에 새 값을 추가
		cap : 새로운 배열의 길이 값
		len : 기존길이에 추가한 개수만큼 더한 값
		포인터 : 새로운 배열을 가르키는 슬라이스 구조체를 return
	*/
}
