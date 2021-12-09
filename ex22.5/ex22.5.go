// Ring : 뱀이 자기 꼬리를 물고 있듯 마지막에서 다시 처음으로 돌아오는 구조
// 링은 언제 쓸까? 일정한 갯수만 사용하고 오래된 요소가 지워져도 되는 경우에 보통 사용
/*
1. 실행 취소 기능(ctrl + z) : 문서 편집기 등에서 일정한 개수의 명령을 저장하고 실행 취소하는 경우
2. 고정 크기 버퍼 기능 : 데이터에 따라 버퍼가 증가되지 않고 고정된 길이로 쓸 때
3. 리플레이 기능 : 게임 등에서 최근 플레이 10초를 다시 리플레이할 때와 같이 고정된 길이의 리플레이 기능을 제공할 때
*/
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
