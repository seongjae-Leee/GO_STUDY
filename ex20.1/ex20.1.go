// 인터페이스 : 구체화된 객체가 아닌 추상화된 상호작용으로 관계를 표현. by the way 메서드는 관계를 나타냄. 이 관계들을 나타낸게 클래스 다이어그램.
// 메서드에는 구현이 포함되어 있다.

//
/*
● 인터페이스 선언 예시

  // 인터페이스이름,인터페이스 선언 키워드
type DuckInterface interface {
	// 내부의 메서드 집합
	Fly()
	Walk(distance int) int
}
-> DuckInterface 라는 관계를 나타내기 위해서는(implement하기 위해서는) Fly와 Walk이라는 두개의 메서드가 필요한 것
*/

/*
● 인터페이스 규칙
1. 메서드는 반드시 메서드명이 있어야 한다.(_ 같은 걸로 무명의 메서드를 만들 수 없다는 뜻)
2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다. 즉 오버로드(같은 함수명, 다른 매개변수 or 리턴)가 안됨
3. 인터페이스에서는 메서드 구현을 포함하지 않는다.
*/

package main

import "fmt"

type Stringer interface {
	String() string
} // String이라는 메서드가 있는 애면 Stringer라는 타입으로 보겠다는 것

type Student struct {
	Name string
	Age  int
}

// Student가 String 이라는 메서드를 가지고 있음. Student가 Stringer 라는 인터페이스를 구현하고 있다
func (s Student) String() string {
	//         Sprintf : 터미널에 출력하는게 아니라 해당 문자열을 만들어서 그걸 반환해주는 것
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
} // 이건 main 함수에서 stringer.getAge() 등 으로 불러올 수 없음

func main() {
	student := Student{"철수", 12}
	var stringer Stringer = student
	// stringer = student .. 위처럼 한줄에 써줘야 함
	fmt.Printf("%s\n", stringer.String())
}

// 결과값 : 안녕! 나는 12살 철수라고 해
