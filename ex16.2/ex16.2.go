// 패키지에는 별칭을 부여할 수도 있다.package main
package main

import (
	"fmt"
	"math/rand"
	_ "strings" // 불러와야 하는 경우, 그러나 쓰지는 않는 경우 별칭을 줘서 쓸 수 있고, _ 를 쓰게 되면 무명으로 불러올 수 있게 된다.
)

func main() {
	//rand.Int() : 유사 랜덤값 반환
	fmt.Println(rand.Int())
}

// 결과값 : 5577006791947779410
