// continue : 후처리로 건너뛴다.
// break : for문을 종료한다.

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // i++(후처리)로 돌아감
		}
		if i == 6 {
			break // 6일 때는 for문 종료
		}
		fmt.Println("6*", i, "=", 6*i)
	}
}
