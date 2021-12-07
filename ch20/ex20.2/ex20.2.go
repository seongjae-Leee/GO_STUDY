package main

// import "goproject/ch20/fedex"
import "goproject/ch20/koreaPost"

// 통합형 인터페이스
type Sender interface {
	Send(parcel string)
}

// // import한 것에서 Send를 호출해서 택배를 보냄
// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }
// func SendBook(name string, sender *koreaPost.PostSender) {
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// sender := &fedex.FedexSender{}
	// sender := &koreaPost.PostSender{}
	var sender Sender = &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}

/*
결과값 :
Fedex sends 어린 왕자 parcel
Fedex sends 그리스인 조르바 parcel
*/
// 근데 이제 한국우체국을 이용하고 싶어졌다면? koreaPost
// 인터페이스를 안 사용하면 샷건 수술이 생기는 거임. 즉 다 바꿔줘야 되는 것
// 그래서 Sender라는 interface를 만들어주면 된다.

// 그렇게 하면 사용하는 부분들은 두고 main에서 호출하는 부분만 바꿔주면 변경사항을 줄일 수 있는 것.

// 즉 내부 동작을 감춰서 서비스 제공자와 사용자 모두에게 자유를 주는 방식!! = 추상화(abstraction)
// 내부 발송 과정은 몰라도 보내는건 문제가 없듯이 추상화를 하게 되면 ★의존성을 끊을 수 있다. low decoupling(프로그램 변경사항 최소화)
