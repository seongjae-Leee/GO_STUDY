// defer 명령문
// -> 지연됐다가 함수 종료 전에 실행을 보장
// 주로 OS자원을 반납해야 할 때 사용한다. 까먹지 않게 처음에 쓰더라도 제일 마지막에 실행이 보장되기 때문이다.
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}
	// defer는 먼저 쓴게 가장 나중에 출력된다.(스택 Stack).. Last In First Out
	defer fmt.Println("반드시 호출된다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")
	fmt.Println("파일에 Hello seongjae를 씁니다.")
	// 출력 string을 정할 수 있게 된다. 즉 test.txt 파일에 적을 내용을 정할 수 있다.
	fmt.Fprintln(f, "Hello seongjae")
}

/*
결과값 :
파일에 Hello seongjae를 씁니다.
파일을 닫았습니다.
반드시 호출된다.
*/
