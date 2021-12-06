package custompkg

import "fmt"

// 아래와 같이 함수 이름이 대문자로 시작하면 다른 파일에서도 import해서 쓸 수 있지만 소문자로 되어있으면 export가 안되어서 불러다 쓸 수 없음
func PrintCustom() {
	fmt.Println("This is custom pakcage~!!")
}

// 아래와 같이 type도 대문자면 외부공개된다. cystompkg.Student{"성재", 27, 100} 이런 식으로...
type Student struct {
	Name  string
	Age   int
	Score int
}

// 아래와 같이 전역변수도 대문자면 외부공개된다. custompkg.Ratio = 10 이런 식으로...
var Ratio int
