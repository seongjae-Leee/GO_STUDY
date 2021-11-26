package main

import "fmt"

func main() {
	fmt.Println("Hello seongjae~")
}

// go run hello.go 를 통해 실행할 수 있고
// cd hello로 와서 go mod init goproject/hello로 모듈을 설치하고 go build로 exe파일을 만들고나서 ./hello.exe로 실행할 수 있다.

/* package main에 빨간줄이 뜨면 설정의 gopls에 들어가서 settings.json에서
"gopls": {
	"experimentalWorkspaceModule": true,
}

를 써주면 된다.
*/
