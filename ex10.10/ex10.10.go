package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a==1")
		// break
		// go에서는 case마다 break를 적지 않아도 case에 해당하면 빠져나가기 때문에 괜찮다.
	case 2:
		fmt.Println("a==2")
	case 3:
		fmt.Println("a==3")
		fallthrough // go에서는 오히려 다른 언어에서처럼 동시에 아래 것도 진행하고 싶을 때 fallthrough로 동시에 진행해준다.
	case 4:
		fmt.Println("a==4")
	default:
		fmt.Println("a != 1, 2, 3, 4, ", a)
	}
}

/*
결과값:
a==3
a==4

이건 fallthrough가 있기 때문에 가능했다.
*/
