# if문

```go
if k == 1 { // 조건식은 반드시 boolean, 구분자로 사용되는 대괄호이기도 하기에 같은 줄에 위치
	println("one")
}

// indent 중요!!!!!!!!!!
if k == 1 {
	println("one")
} else if k == 2 {
	println("two")
} else {
	println("other")
}

// Optional Statement
// 조건식 이전에 실행 가능 but if문 안에서만 적용
if val := i * 2; val < max {
	println(val)
}
```

# Switch문

```go
package main

func main(){
  var name string
  var category = 1

  switch category {
  case 1:
    name = "Paper"
  case 2:
    name = "eBook"
  case 3, 4:
    name = "Blog"
  default:
    name = "Other"
  }
  println(name)


  // Expression 사용
  // expression : switch문 내부에서 사용할 변수
  switch x := category << 2; x - 1 {
    //.......
  }
}


// switch 뒤에 조건변수 or expression 쓰지 않아도 됨
// 이럴 경우 case문을 순서대로 검사하여 조건에 맞는 경우 해당 case 실행
func grade(score int) {
    switch {
    case score >= 90:
        println("A")
    case score >= 80:
        println("B")
    case score >= 70:
        println("C")
    case score >= 60:
        println("D")
    default:
        println("No Hope")
    }
}   

// 변수 v의 data 타입을 확인 후 타입에 맞는 case 실행
switch v.(type) {
case int:
    println("int")
case bool:
    println("bool")
case string:
    println("string")
default:
    println("unknown")
}  


// break 필요 x -> break를 적든 안적든 항상 break 
// Go에서 break하지 않고 아래 명령어를 수행하기 위해서 fallthrough를 사용
package main

import "fmt"

func main() {
    check(2)
}

func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}

```

-	다른 언어는 switch 키워드 뒤에 변수나 expression 반드시 두지만, Go는 이를 쓰지 않아도 된다. 이 경우 Go는 switch expression을 true로 생각하고 첫번째 case문으로 이동하여 검사
-	다른 언어의 case문은 일반적으로 리터럴 값만을 갖지만, Go는 case문에 복잡한 expression 가능
-	다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동 but Go 는 break 필요 x
-	다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, Go는 그 변수의 Type에 따라 case로 분기 가능


