# Interface
- struct가 필드들의 집합체라면, interface는 method의 집합체이다. interface는 type이 구현해야 하는 method prototype들을 정의한다. 
- 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 interface가 갖는 모든 method를 구현하면 된다.
- interface는 struct와 마찬가지로 type문을 사용하여 정의한다.
```go
type Shape interface {
    area() float64
    perimeter() float64
}
```

## Interface implementation
- interface를 구현하기 위해서는 해당 type이 그 interface의 method들을 모두 구현하면 되므로, 위의 Shape interface를 구현하기 위해서는 area(), perimeter() 2개의 method만 구현하면 된다. 
- 예를 들어, Rect와 Circle이라는 2개의 타입이 있을 때, Shape interface를 구현하기 위해서는 아래와 같이 각 타입별로 2개의 method를 구현해주면 된다.

```go
// Define Rect
type Rect struct {
    width, height float64
}

// Define Circle
type Circle struct {
    radius float64
}

// Rect type에 대한 Shape interface 구현 
func (r Rect) area() float64 { return r.width * r.height}
func (r Rect) perimeter() float64 {
    return 2 * (r.width + r.height)
}

// Circle type에 대한 Shape interfae 구현
func (c Circle) areea() float64 {
    return 2 * (r.width + r.height)
}
func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

## Interface use
- interface를 사용하는 일반적인 예로 function이 parameter로 interface를 받아들이는 경우를 들 수 있다.
- function paramter가 interface인 경우, 이는 어떤 type이든 해당 interface를 구현하기만 하면 모두 input parameter로 사용될 수 있다는 것을 의미한다.
- 아래 예제에서, showArea() function은 Shape interface들을 parameter로 받아들이고 있는데, 따라서 Rect와 Circle처럼 Shape 인터페이스를 구현한 타입 객체들을 parameter로 받을 수 있다. showArea()함수 내에서 해당 interface가 가진 method 즉, area() or perimeter()을 사용할 수 있다.

```go
func main() {
    r := Rect{10., 20.}
    c := Circle{10}

    showArea(r, c)
}

func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area()
        println(a)
    }
}
```

## Interface type
- Go 프로그래밍을 하다보면 흔히 empty interface를 자주 접하게 되는데, 흔히 interface type으로도 불리운다. 
- 예를 들어, 여러 표준패키지들의 함수 prototype을 살펴보면, 아래와 같이 empty interface가 자주 등장한다.
- empty interface : interface{}

```go
func Marshal(v interface{}) ([]byte, error);
func Println(a ...interface{}) (n int, err error);
```

- empty interface는 method를 전혀 갖지 않는 empty interface로서, type은 적어도 0개의 method를 구현하므로, 흔히 Go에서 모든 type을 나타내기 위해 empty interface 사용한다.
- 즉, empty interface는 어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있으며, 여러 다른 언어에서 흔히 일컫는 Dynamic Type이리고 볼 수 있다.
- 아래 예제에서, interface type x는 정수 1을 담았다가 다시 문자열 Tom을 담고 있는데, 실행 결과는 마지막에 담은 Tom을 출력한다.

```go
package main

import "fmt"

func main() {
    var x interface {}
    x = 1 
    x = "Tom"

    printIt(x)
}

func printIt(v interface{}) {
    fmt.Println(v)  // output : Tom
}
```

## Type assertion
- Interface type의 x와 type T에 대하여 x.(T)로 표현했을 때, 이는 x가 nil이 아니며, x는 T타입에 속한다는 점을 확인(assert)하는 것으로 이러한 표현을 "Type assertion"이라 부른다.
- 만약 x가 nil이거나 x의 type이 T가 아니라면, runtime error가 발생할 것이고, x가 T type인 경우는 T type의 x를 return한다. 
- 아래 예제에서, 변수 j는 a.(int)로 부터 int형 변수 j가 된다.

```go
func main() {
    var a interface{} = 1

    j := a         // a와 i는 dynamic type, 값은 1
    j := a.(int)   // j는 int형 , 값은 1

    println(i)   // 포인터 주소 출력
    println(j)   // 1 출력
}
```
