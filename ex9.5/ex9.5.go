package main

import "fmt"

func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("우리끼리 나눠내자")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println("아이고 신발끈이")
		} else {
			fmt.Println("아이고 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 낼게")
	}
}
