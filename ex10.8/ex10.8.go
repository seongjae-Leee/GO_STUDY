package main

import "fmt"

// 새로운 타입을 만드는 것 (그냥 int라고 안 쓰고 만들어준 이유는 아래 색깔들의 분류와 용도를 명확히 하기 위함.)
type ColorType int

// 0부터 1씩 증가하는 정수로 만들어두고(javascript의 배열처럼)
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	// Blue가 출력값으로 나오는 함수를 만들어준 것.
	return Blue
}

func main() {
	fmt.Println("My Favorite color is", colorToString(getMyFavoriteColor()))
}

// 결과값은 My Favorite color is Blue
