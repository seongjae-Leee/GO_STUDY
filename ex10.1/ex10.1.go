// switch 문에서는 위에서부터 탈출구를 찾으며 내려옴. 탈출못하면 마지막에 default 도달
package ex101

import "fmt"

func main() {
	a := 3

	// swtich의 값과 case의 조건을 비교한다. case의 조건에 콤마(,)로 구분된게 있으면 || 연산으로써 둘중 하나만 true여도 됨.
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fmt.Println("a == 3")
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1, 2, 3", a)
	}
}

// if와 else if 떡칠보다는 하나의 값과 비교할 경우에는 switch 구문 쓰는게 가독성이 좋은듯.
