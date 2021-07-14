# Method
Go는 객체지향 프로그래밍(OOP)을 고유의 방식으로 지원한다.
method는 특별한 형태의 func이다. method는 함수 정의에서 func 키워드와 함수명 사이에 "그 
함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다. 흔히 receiver로 불리우는 이 부분은 method가 속한 struct타입고ㅘ struct 변수명을 지정하는데, struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다. 예를 들어, 아래 예제는 Rect라는 struct를 정의하고 area()라는 method를 정의하고 있다. func와 area() 사이에 Rect 타입의 r이 정의되고 이를 함수 본문에서 사용하고 있다. method가 선언된 이후에는 Rect 구조체의 객체는 rect.area() 문장처럼 area() method를 struct 객체로 부터 직접 호출할 수 있다. 

```go
package main

//Rect - struct 정의
type Rect struct {
    width, height int
}

//Rect의 area() method
func (r Rect) area() int {
    return r.width * r.height
}

func main() {
    rect := Rect{10, 20}
    area := rect.area()
    println(area)
}
```

## Value vs pointer receiver
value receiver는 struct의 data를 copy하여 전달, pointer receiver는 struct의 pointer만을 전달 value receiver의 경우 만약 method내에서 그 struct의 필드값이 변경되더라도 호출자의 데이터는 변경되지 않는 반면, pointer receiver는 method 내의 필드값 변경이 그대로 호출자에서 반영<br/>

위의 Rect.area() method는 value receiver를 표현한 것으로 만약 area() method 내에서 width나 height가 변경되더라도 main() 함수의 rect 구조체의 필드값에는 변화가 없다. 하지만, 아래와 같이 이를 pointer receiver로 변경한다면, method 내 r.width++ 필드 변경분이 ain() 함수에서도 반영되기 때문에 output : 11, 220 

```go
// pointer receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}

func main() {
    rect := Rect{10, 20}
    area := rect.area2()
    println(rect.width, area)s
}