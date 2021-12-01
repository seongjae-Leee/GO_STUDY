/* for문 기본 구문
for 초기문; 조건문; 후처리{
	코드 블록    <- 조건문이 true일 때 실행됨.
}
*/
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
}
