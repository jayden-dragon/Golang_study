# Go Error
- Go는 내장 타입으로 error라는 interface를 갖는다.
- Go error는 이 error interface를 통해서 주고 받게 되는데, 이 interface는 다음과 같은 하나의 method를 갖는다. 
- 개발자는 이 인터페이스를 구현하는 커스텀 error type을 만들 수 있다.
```go
type error interface {
    Error() string
}
```

## error
- Go 함수가 결과와 error를 함께 return한다면, 이 error가 nil인지를 체크해서 