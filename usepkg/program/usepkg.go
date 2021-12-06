package main

import (
	"fmt"
	"goproject/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 5}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

// go mod tidy로 import한 모듈 설치 -> go.sum에서 다운받은 패키지들을 볼 수 있음
// go env -> go environment의 약자로 go의 설정들을 볼 수 있음 -> GOPATH 확인해보면 C:\Users\user\go에 있음
// C:\Users\user\go 에서 ls 해보면 bin과 pkg를 볼 수 있음.
// cd pkg 해보면 mod하고 sumdb가 있음
// cd mod 해보면 우리가 다운받은 패키지가 있음
// 쨌든 usepkg/program 가서 go build
// 하고서 ./program.exe 실행

// 실행하게 되면 usepkg 의 main함수가 실행되므로 custompkg.PrintCustom()이 먼저 실행됨...
