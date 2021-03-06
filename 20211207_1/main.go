package main

func main() {
	// interface : 구현을 포함하지 않는 메서드의 집합
	// 장점 : 구체화된 타입이 아닌 추상적인 개념. 따라서 프로그램 요구사항 변경시 유연하게 대처 가능.
	// Go에서는 인터페이스 구현 여부를 그 타입이 인터페이스에 해당하는 메서드를 가지고 있는지로 판단하는 덕 타이핑을 지원한다.

	// 타입 키워드 써주고 인터페이스 이름 써주고 인터페이스로 쓰겠다는 키워드를 써주면 됨.
	type DuckInterface interface { // interface도 타입중 하나이기 때문에 인터페이스 변수 선언이 가능하고 변수의 값으로도 쓸 수 있다.
		// ▼ 메서드의 집합
		// 주의할 점 : 메서드는 반드시 이름이 있어야 함. 매개변수와 리턴이 다르더라도 이름이 같은 메서드는 안된다.
		// 또한 인터페이스에서는 메서드 구현을 포함하지 않는다.
		walk()
		Move(distance int) int
	}
}
