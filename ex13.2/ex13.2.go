package main

import "fmt"

type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	UserInfo User // 여기서 UserInfo 라는 요소를 만들지 말고 그냥 User라고만 썼을 경우 아래 출력하는 부분에서 vip.UserInfo.로 들어가지 말고 바로 User.으로 갈 수 있다.
	VIPLevel int
	Price    int
}

func main() {
	user := User{"이성재", "seongjae123", 27}
	vip := VIPUser{
		User{"홍길동", "gildonggg", 43},
		3,
		999999999,
	}
	fmt.Printf("유저 : %s ID : %s 나이 : %d\n", user.Name, user.Id, user.Age)
	fmt.Printf("VIP 유저 : %s ID : %s 나이 : %d VIP레벨 : %d, 소비가격 : %d\n",
		vip.UserInfo.Name, vip.UserInfo.Id, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}
