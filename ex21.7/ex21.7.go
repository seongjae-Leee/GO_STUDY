// ● 의존성 주입 : 외부에서 로직을 주입하는 것
package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// 네트워크랑 상관없이 로직이 분리되어 실행됨.
func writeHello(writer Writer) {
	writer("Hello seongjae")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		//  파일에 새기기
		fmt.Fprintln(f, msg)
	})
}
