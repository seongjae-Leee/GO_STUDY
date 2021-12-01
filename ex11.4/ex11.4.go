package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// standard input
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("뭐라도 입력하세요")
		var number int
		// _, 는 빈칸지시자로써, 첫번째 리턴값은 안 쓰겠다는 뜻
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자로 입력하세요")
			// 키보드 버퍼를 지운다.
			stdin.ReadString('\n') // <- 개행문자가 나올 때까지 키보드 버퍼를 읽어온다.
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break // 짝수면 break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}

/*
결과값 :

뭐라도 입력하세요
ld
숫자로 입력하세요
뭐라도 입력하세요
13
입력하신 숫자는 13입니다.
뭐라도 입력하세요
14
입력하신 숫자는 14입니다.
for문이 종료되었습니다.
*/
