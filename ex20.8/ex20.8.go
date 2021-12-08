package main

type Stringer interface{
	String() string
}

type Student struct{
	//
}

func main(){
	var stringer Stringer
	stringer.(*Student) // 컴파일 에러 등장
	// impossible type assertion. -> stringer를 Student의 포인터로 타입 변환이 불가능하다는 것이다.
	// *Student does not implement Stringer(missing String method)