// 타입변환 주의사항
package main

import "fmt"

func main() {
	var a int16 = 3456   //a는 int16타입으로 2바이트 정수
	var b int8 = int8(a) // int16->int8

	// a는 3456인데 b는 -128이 나옴. 큰 값을 작은 값으로 옮겼기 때문!!
	fmt.Println(a, b)
}
