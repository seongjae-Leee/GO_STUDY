// 패키지에는 별칭을 부여할 수도 있다.package main
package main

import (
	"fmt"
	"math/rand"
	_ "strings" // 불러와야 하는 경우, 그러나 쓰지는 않는 경우 별칭을 줘서 쓸 수 있고, _ 를 쓰게 되면 무명으로 불러올 수 있게 된다.
	// ★왜 사용하지 않는 패키지를 포함해야 하는가? -> 패키지 초기화에 따른 부가효과를 위해서.
)

func main() {
	//rand.Int() : 유사 랜덤값 반환
	fmt.Println(rand.Int())
}

// 결과값 : 5577006791947779410

/*
Go 모듈 만들고 외부 패키지 활용하기
1. goproject/usepkg 폴더 생성
2. go mod init goproject/usepkg
3. mkdir custompkg
4. custompkg.go
5. mkdir program
6. usepkg.go
7. go mod tidy
*/
