// 패키지 : 코드를 묶는 단위. 모든 코드는 반드시 패키지로 묶여야 한다.
// 프로그램 : 실행 시작 지점을 포함한 패키지. 즉 main 함수를 포함한 main 패키지
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Int() : 유사 랜덤값 반환
	fmt.Println(rand.Int())
}

// 결과값 : 5577006791947779410
