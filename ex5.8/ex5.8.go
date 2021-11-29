package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//                            ▼ standard input
	stdin := bufio.NewReader(os.Stdin)
	// var stdin = bufio.NewReader() 이거랑 같은 거임

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	//         ▼ 아무것도 아닌 메모리값(null과 유사한 개념)
	if err != nil {
		fmt.Println(err)
		// ▼ '표준입력'에서 '문자를 읽어오라' '\n이 나올 때까지' ...cf) \n은 개행문자
		stdin.ReadString('\n')
	} else {
		fmt.Println(a, b)
	}

	// 위에서 입력을 받아왔으므로 그냥 쓰면 된다. := 은 선언을 해주는 거기 때문에 그렇게 하지 말고 =로. 혹은 아예 지우거나.
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}

// # goproject/ex5.8
// .\ex5.8.go:17:2: n declared but not used  라는 에러 발생
// 즉 n이 쓰이지 않아서 그런 것이므로 33번째 줄에 그냥 하나 넣어주자.
// go에서는 쓰이지 않는 것은 선언조차 할 필요가 없다.

/*
PS C:\Users\user\goproject\ex5.8> ./ex5.8.exe
hello 4
expected integer (err)
10 4
2 10 4
*/
