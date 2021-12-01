package main

import "fmt"

type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	UserInfo User
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
