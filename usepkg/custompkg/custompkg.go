package custompkg

import "fmt"

// 아래와 같이 함수 이름이 대문자로 시작하면 다른 파일에서도 import해서 쓸 수 있지만 소문자로 되어있으면 export가 안되어서 불러다 쓸 수 없음
func PrintCustom() {
	fmt.Println("This is custom pakcage~!!")
}
