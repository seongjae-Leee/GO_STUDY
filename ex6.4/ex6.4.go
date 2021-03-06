package main

import "fmt"

func main() {
	var x int8 = 127
	// 정수로  //불리언값
	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1)
	// 돌려보면 127 < 127+1 이면 true 여야 하는데 false가 나옴. 왜 그런지 알아보자.
	//          탭, 4자리로 출력(정수 한번, 이진수 한번)
	fmt.Printf("x\t=%4d, %08b\n", x, x)
	fmt.Printf("x+1\t=%4d, %08b\n", x+1, uint8(x+1))
	// go build하고 다시 돌려보면 127과 -128이 나온다.
	// 왜 -128이 되냐 하면 8자리에서는 부호비트, 즉 맨 앞에 있는 숫자가 0이면 양수 1이면 음수로 인식하기 때문이다. 즉 1바이트 정수에서는 -128~127까지 나타낼 수 있는데
	// 127이 가장 큰 수이기 때문에 128은 음수로, 즉 숫자가 overflow 난 것이다.
}
