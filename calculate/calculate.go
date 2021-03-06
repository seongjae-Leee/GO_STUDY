package main

import "fmt"

func main() {
	a := 2
	var num int //0

	num = a
	fmt.Println("num = a :", num) // 2
	num += 4
	fmt.Println("num += 4 :", num) // 6
	num -= 2
	fmt.Println("num -= 2 :", num) // 4
	num *= 5
	fmt.Println("num *= 5 :", num) // 20
	num /= 2
	fmt.Println("num /= 2 :", num) // 10
	num %= 3
	fmt.Println("num %= 3 :", num) // 1

	num = 3                                       //0000 0011
	num &= 2                                      //0000 0010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num) //0000 0010 = 2
	num |= 5                                      //0000 0101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num) //0000 0111 = 7
	// 이 아래가 헷갈림...
	num ^= 4                                       //0000 0100
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num)  //0000 0011 = 3
	num &^= 2                                      //0000 0010
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num) //0000 0001 = 1
	num <<= 9                                      //0000 1001
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num) //1000 0000 = 512
	num >>= 8                                      //0000 1000
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num) //0000 0010 = 2
}
