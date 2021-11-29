package main

import (
	"fmt"
	"math"
)

// 둘이 같은지 비교해서 boolean 값으로 결과를 내주는 함수
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	//         소수점 18자리까지
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))
	// 0.299999999999999989 == 0.300000000000000044 : true
	// 왜 같지 않은 값인데 true가 나올까?
	// equal 함수에서 쓰인 math.Nextafter 메서드가 1비트만큼 a+b를 c를 향해 움직여주기 떄문이다.

	fmt.Printf("%0.18f == %0.18f : %v\n", a+b, c, equal(c, a+b))
	// 순서를 바꿔봐도 마찬가지다. 결국 옮겨준다는 개념보다는 둘이 1bit 차이인지를 '비교'해주는 메서드라고 보면 될 것 같다.
}
