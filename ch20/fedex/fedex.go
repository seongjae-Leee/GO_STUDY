// 인터페이스 왜 쓰나?

package fedex

import "fmt"

type FedexSender struct {
	// ...
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %s parcel\n", parcel)
} // 택배 발송
