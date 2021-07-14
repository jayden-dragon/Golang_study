# Function

`func 키워드() {}` : Go에서는 다음과 같은 방식으로 함수를 선언, 함수가 반드시 앞에 선언될 필요 x
```go
package main

func main() {
    msg := "hello"
    say(msg)
}

func say(msg string) {
    println(msg)
}
```

## Pass By Value vs Pass By Reference

파라미터 전달 방식

1. Pass By Value : 값을 전달, main문 msg != say함수 msg
2. Pass By Reference : 값의 주소를 전달, main문 msg == say함수 msg

```go
package main 
func main() {
    msg := "hello"
    say(&msg)
    println(msg)
}

func say(msg *string) { 
    println(*msg)
    *msg = "Changed"    // Dereference
}

/* Output
hello
Changed
*/
```

## Variadic Function (가변인자함수)
함수에 고정된 숫자가 아니라 다양한 숫자의 파라미터를 전달하고자 할 때 사용 <br/>
여러개의 동일한 타입의 파라미터 전달 가능 

```go
package main
func main() {
    say("This", "is", "sparta")
    say("Hi")
}

func say(msg ...string) {
    for  _, s := range msg {
        println(s)
    }
}

/* Output
This
is
sparta
Hi
*/
```

## Return value
리턴값이 있을 수도 없을 수도, 하나 일 수도, 복수 일 수도 있다.

*Named Return Parameter* : 리턴되는 값을 (함수에 정의된) 리턴 파라미터들에 할당 가능
`func sum(nums ...int) int {}` : 리턴되는 데이터 타입을 파라미터 뒤에 위치

```go
// 리턴값 1개
package main

func main() {
    total := sum(1, 7, 3, 5, 9)
    println(total)
}

func sum(nums ...int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    return s
}

// 리턴값 여러 개
package main

func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)
}

func sum(nums ...int) (int, int) {
    s := 0
    count := 0
    for _, n := range nums {
        s += n
        count ++
    }

    return count, s 
}

// 리턴값 여러개 가독성 up
func sum(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count =len(nums)
}
return
```


## Anonymous Function
일반적으로 익명함수는 그 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어 사용된다. <br/>
함수명이 없고 변수에 할당되어 사용되기 때문에 변수명이 함수명처럼 사용된다.

```go
package main
func main() {
    sum := func(n ...int) int {
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }

    result := sum(1, 2, 3, 4, 5)
    println(result)
}
```

## First Class Function

```go
package main
func main() {
    add := func(i int, j int ) int {
        return i + j
    }

    r1 := calc(add, 10, 20)
    println(r1)

    r2 := calc(func(x int, y int) int { return x - y}, 10, 20)
    println(r2)
    
}

func calc(f func(int, int) int, a int, b int) int {
    result := f(a, b)
    return result
}
```

## type문을 사용한 함수 원형 정의

```go
type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```