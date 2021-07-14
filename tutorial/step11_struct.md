# Struct
- Go에서 struct는 custom data type를 표현하는데 사용
- Go의 struct는 필드들의 집합체이며 필드들의 컨테이너
- Go에서 struct는 필드 데이터만을 가지며, method는 갖지 x

## Struct declare
- Custom type을 정의하기 위해 type문을 사용

```go
package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    // person 객체 생성
    p := person{}
    
    // 필드값 설정
    p.name = "Lee"
    p.age = 10

    fmt.Println(p)
}
```

## Struct 객체 생성

```go
var p1 person 
p1 = person{"Bob", 20}
p2 := person{name: "Sean", age: 50}

p := new(person)
p.name = "Lee"
```

## Constructor
map 형태의 구조체의 경우 초기화가 필요. 이를 위해 생성자 함수를 통해 구조체를 직접 호출하지 않고 생성자 함수를 호출해 map 초기화 후 선언

```go
package main

type dict struct {
    data map[int]string
}

func newDict() *dict {
    d := dict{}
    d.data = map[int]string{}
    return &d // 포인터를 전달
}

func main(){
    dic := newDict()
    dic.data[1] = "A"
}
```

