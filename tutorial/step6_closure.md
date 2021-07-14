# Closure
함수 바깥에 있는 변수를 참조하는 함수값을 일컫는데, 이때 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.

외부 함수를 약간 익명함수처럼 사용하는 느낌 

```go
package main

func nextValue() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    next := nextValue()

    println(next())
    println(next())
    println(next())

    anotherNext := nextValue()
    println(anotherNext())
    println(anotherNext())
}

/* 출력
1
2
3
1
2
*/
```

