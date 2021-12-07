// 덕 타이핑. 즉 이용자는 자기가 보이는 것이 전부다. UI도 마찬가지인 것 같다. 이용자는 추상 계층에 관심이 없기 때문에
// 그리고 프론트와 백도 서로에 대해 잘 몰라도 중간 추상계층을 확실히 만들어놓아야 진정한 UI를 이룰 수 있다.
// 내가 생각하는 진정한 UI란 이용자가 편리하게 이용할 수 있게 하는 것 뿐 아니라 그에 앞서 우리가 전달하고자 하는 바를
// 이용자가 의식하지 않고도 정확히 전달받아야 하는 것이다.

// 인터페이스 구현 여부를 타입 선언시 하지 않고, 인터페이스가 사용될 때 결정하기 때문에 서비스 제공자는 구체화된 객체를 제공하고 사용자가
// 필요에 따라 인터페이스를 정의해서 사용할 수 있는 사용자 중심의 코딩이 된다.
// 인터페이스는 필요에 의해 구현하는게 좋다. 미리 예상하고 다 구현해 놓는 것은 낭비가 될 수 있기 때문이다.

// 그리고 interface는 실해중(runtime)에 검사하는 것이 아니라 컴파일 될 때 검사를 하기 때문에 성능상 문제가 없고
// Go언어 자체가 빌드 시간이 워낙 빠르기 때문에 컴파일할 때도 딱히 문제는 없다.