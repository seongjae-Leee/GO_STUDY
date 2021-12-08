package main

import "fmt"

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
	//...
}

func (c CDatabase) GetData() {
	//
}

func (c CDatabase) SetData() {
	//
}

// 위 두개는 Database interface에 없는 메서드를 사용중인데 이 때 함부로 Database interface의 메서드를 바꿀 수는 없다.
// 그럼 어떻게 해야 하나
type Wrapper struct {
	cdb CDatabase
}

func (w Wrapper) Get() {
	w.cdb.GetData()
}
func (w Wrapper) Set() {
	w.cdb.SetData()
}

// 위처럼 만들어 놓으면 아래처럼 쓸 수 있다
func main() {
	var cdatabse CDatabase
	var database Database
	database = Wrapper{cdatabse}
	fmt.Println(database)
}
