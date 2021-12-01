// 무한 루프
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		// 1초동안 프로그램을 멈춰라
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}

/*
결과값
1
2
3
4
5
6
...
1초마다 1씩 증가하면서 계속 숫자가 나옴
*/
