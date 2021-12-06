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

	slice3 := []int{1, 2, 3}       // len : 3, cap : 3
	slice4 := append(slice3, 4, 5) // 공간이 부족해서 2배로 늘린 6개짜리 capability의 배열을 만들어서 slice 배열의 모든 값을 복사하고 추가해줌.
	// 즉 len : 5, cap : 6
	fmt.Println("슬라이스3 : ", slice3, len(slice3), cap(slice3))
	fmt.Println("슬라이스4 : ", slice4, len(slice4), cap(slice4))

	slice3[1] = 200
	fmt.Println("슬라이스3 : ", slice3, len(slice3), cap(slice3))

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/*
		슬라이싱 : 배열의 일부를 집어낸다..
		array[startIndex : endIndex]
		주의할 점 : endIndex-1까지
	*/
	arr := [5]int{1, 2, 3, 4, 5}
	slice5 := arr[1:3] // 슬라이싱(index 1부터 (3-1)까지)

	fmt.Println("arr : ", arr)
	fmt.Println("slice5 : ", slice5, len(slice5), cap(slice5))
	/*
		arr :  [1 2 3 4 5]
		slice5 :  [2 3] 2 4
	*/
	// arr의 두번째 값을 변경해보자
	arr[1] = 100

	fmt.Println("arr : ", arr)
	fmt.Println("slice5 : ", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 200)
	fmt.Println("arr : ", arr)
	fmt.Println("slice5 : ", slice5, len(slice5), cap(slice5))
	/*
		arr :  [1 100 3 4 5]
		slice5 :  [100 3] 2 4
		arr :  [1 100 3 200 5]
		slice5 :  [100 3 200] 3 4
	*/

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 슬라이스를 슬라이싱
	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := slice6[1:3] // [2 3]
	fmt.Println(slice7)
	// 처음부터 하고 싶으면 slice6[:3] 즉 0 생략 가능
	// 끝까지 하고 싶으면 slice6[1:len(slice6)]
	// 전체를 하고 싶으면 slice6[:]  -> 배열 전체를 가르킬 때 주로 씀
}
