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

// 참고 : 논리 연산자
/*
&& = 양변이 모두 true일 때 true 반환
|| = 양변중 하나라도 true라면 true 반환
! = true이면 false를 반환하고 false이면 true 반환.. ex) !(2<5 || 10<5) = false. 왜냐하면 둘중 하나가 참이라서 true지만 !true 이므로 false
*/

// 참고 : 대입 연산자
/*
a, b = a', b' 일때 a' 값이 a에 들어가고 b' 값이 b에 들어간다.
*/

// 참고 : 복합 대입 연산자
/*
a = 10일 때
a +=2 라면 a에 2를 더해줘야 하므로 a = a+2 = 12
... -, *, /, %도 마찬가지
*/

// 참고 : 증감 연산자
/*
a = 10일 때
a++ 이면 a값에 1 더해줘라
a-- 이면 a값에 1 빼줘라 이런 말
대신!! 값을 반환하지는 않음.
*/
