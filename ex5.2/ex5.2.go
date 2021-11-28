package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a : ", a, "b : ", b)
	fmt.Println("a : ", a, "b : ", b, "f : ", f)
	// Printf는 formatter가 필요. %d 는 decimal이라는 정수 포맷으로 나타낸다는 뜻. %f는 실수를 나타내고
	fmt.Printf("a : %d b : %d f : %f\n", a, b, f)
}

/*
-출력결과 (Print와 달리 Println은 콤마를 기준으로 한칸씩 띄워서 보여준다.)
a : 10 b : 20     a :  10 b :  20 f :  3.27994387438297e+10(지수 표현으로 해당 숫자에 10^10을 곱해주어야 한다는 뜻이다.)
a : 10 b : 20 f : 32799438743.829700
*/
