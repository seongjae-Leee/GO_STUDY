package exinit

import "fmt"

var (
	a = c + b //c 값이 먼저 정해져야함
	b = f()   // 아래 c와 마찬가지로 5
	c = f()   // 그럼f 함수로 가야함. 갔다 오면 4
	d = 3
)

// 패키지 초기화. 이 초기화 이후 export되는 것. 여기서 중요한 점!! ★이 패키지가 custompkg에도 import되어도 init 함수 자체는 한번만 실행된다.
// 그리고 custompkg에 import 되었다면 이 exinit 패키지에서는 custompkg를 import 할 수 없다. cycle 오류
func init() {
	d++
	fmt.Println("exinit.init function", d)
}

func f() int {
	d++
	fmt.Println("f() d : ", d)
	return d
}

func PrintD() {
	fmt.Println("d : ", d)
}
