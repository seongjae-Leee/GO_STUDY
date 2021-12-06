// 동적 배열 & 가변 배열 (자료구조).. data를 얼마나 효율적으로 관리를 할 것이냐.
package main

import "fmt"

func main() {
	// slice : 동적 배열(자동으로 배열 크기를 증가시켜 준다.)
	// 장점 : 길이가 요소 개수에 따라 자동으로 증가해 관리가 편함.
	// 슬라이싱 기능을 사용해서 배열의 일부를 나타내는 슬라이스를 만들 수 있다.

	// var arr [10]int -> 10개까지..

	var slice []int

	if len(slice) == 0 {
		fmt.Println("슬라이스가 비어있다.", slice)
	}

	// slice[1] = 10
	// 할당되지 않은 메모리공간에 접근해서 생기는 에러 panic 발생
	// 따라서!!!! 아래와 같이 객체화해서 초기화 필요
	// var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3} // 0번째 index에 1, 5번째 index에 2, 10번째 index에 3 넣기
	fmt.Println(slice2)

	// cf)헷갈리면 안 되는게 var arr = [...]int{1, 2, 3}  이건 슬라이스가 아니라 배열임.

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// make()를 이용한 초기화
	// var slice3 = make([]int, 3) // 첫번째 인수로 타입을 넣어주고, 두번째 인수로 길이

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// 두번째 인덱스를 5로 변경하겠다는 배열식 접근
	var slice4 = make([]int, 3)
	slice4[1] = 5 //[0, 5, 0]

	// 슬라이스 순회 -> 동적으로 길이가 늘어나느거 말고는 배열과 동일함.
	// for문을 통한 초기화
	var slice5 = make([]int, 3)
	for i := 0; i < len(slice5); i++ {
		slice5[i] = i + 1
	}
	// 출력
	for _, v := range slice5 {
		fmt.Println(v)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// 요소 추가 using append
	var slice6 = []int{1, 2, 3}
	slice7 := append(slice6, 4, 5, 6)
	fmt.Println(slice6)
	fmt.Println(slice7)
	// append는 : 첫번째 인수로 들어온 슬라이스의 값을 변경하는게 아니라 요소가 추가된 새로운 슬라이스를 반환
	// 따라서 기존 슬라이스에 요소를 추가하고 싶다? append() 결과를 기존 슬라이스에 대입해서 변경해야 한다.
	// for문으로 한번 추가해보자

	var slice8 []int
	for i := 0; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	slice8 = append(slice8, 11, 3, 316, 123)
	fmt.Println(slice8)

	// type SliceHeader struct {
	// 	Data uintptr // 실제 배열을 가르키는 포인터
	// 	Len  int     // 요소 개수
	// 	Cap  int     // 실제 배열의 길이
	// } -> 슬라이스의 구조

}
