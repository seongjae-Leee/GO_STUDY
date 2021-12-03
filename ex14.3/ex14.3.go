// var p *int
// if p != nil {
// 	//p가 nil이 아니라는 얘기는 p가 유효한 메모리 주소를 가르킨다는 뜻이다.
// }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// package main

// import "fmt"

// type Data struct {
// 	value int      //8바이트
// 	data  [200]int //1600바이트
// }

// //parameter 이름은 arg이고 taype은 Data인 것
// func ChangeData(arg Data) {
// 	arg.value = 999
// 	arg.data[100] = 999
// }

// func main() {
// 	var data Data // 새로 정의되었음. 즉 value의 사이즈도 0이고 data[100]도 사이즈가 0이다.
// 	ChangeData(data)
// 	fmt.Printf("value = %d\n", data.value)
// 	fmt.Printf("data[100] = %d\n", data.data[100])
// }

/*
결과값 :
value = 0
data[100] = 0
*/

// 중요한건 뭐냐면 = 을 통해 우변값을 좌변에 할당할 때 그 '사이즈'가 고대로 할당된다는 것!

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 자 위에 걸 그럼 999, 999가 결과값으로 나오게 하려면 어떻게 바꾸느냐
package main

import "fmt"

type Data struct {
	value int      //8바이트
	data  [200]int //1600바이트
}

// *Data로 포인터로 받으면서 값을 넘기는게 아니라 주소값을 넘기게 하면 됨
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data
	ChangeData(&data) // <- 주소값을 받아옴(Data의 포인터 타입, 즉 *Data 타입)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

/*
결과값 :
value = 999
data[100] = 999
*/
