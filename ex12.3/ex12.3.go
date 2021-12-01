package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // =[5]int{10, 20 ...}
	nums[2] = 3000
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	// // 현석이꺼
	// var abcd []int
	// abcd = []int{1, 2, 3, 4}
	// fmt.Println(len(abcd))
}
