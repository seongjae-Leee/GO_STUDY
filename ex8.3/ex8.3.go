// ★비트클리어 연산을 통해 rooms로 변수를 하나만 쓸 수 있으므로 메모리를 적게 사용할 수 있다. 또한 덜 귀찮다.
package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room // | 이라는 OR 연산자를 통해 비트 플래그를 올려준다 (0 -> 1)
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room // &^ 이라는 비트클리어 연산자를 통해 비트 플래그를 내려준다. (1 -> 0)
} // 비트클리어에 대해서 좀 더 자세히 말해보자면 rooms의 0000 0111과 ^연산자로 연산한 LivingRoom의 0000 0010을 뒤집은 1111 1101을 & 연산해서
// 0000 0101이 되는 것이다.

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
} // 불이 켜져있는지 확인

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다.")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다.")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다.")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다.")
	}
}

func main() {
	var rooms uint8 = 0
	// 어떤 방에 불을 켤건지 세팅을 해놓고(비트 플래그를 이용해 MasterRoom, LivingRoom의 첫번째 비트를 0에서 1로 바꾼 것)
	rooms = SetLight(rooms, MasterRoom) // 0000 0001 (위에 iota로 정해줬기 때문에)
	rooms = SetLight(rooms, LivingRoom) // 0000 0011
	// rooms = SetLight(rooms, BathRoom) ... 화장실은 빼놓고(주석으로) 만약 넣었다면 0000 0111
	rooms = SetLight(rooms, SmallRoom)

	// 작은방은 안 켤래
	rooms = ResetLight(rooms, SmallRoom) // 1을 0으로 바꿔주는 것
	// 이제 켜보자
	TurnLights(rooms)
}
