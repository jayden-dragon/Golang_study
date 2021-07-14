# For문

`for 초기값; 조건; 증감식 {}` : 기본 구조
```go
package main

func main(){
    sum := 0 
    for i := 1; i <= 100; i++ {
        sum += i
    }
    println(sum)
}
```

`for 조건 {}` : while문과 같음

```go
package main

func main(){
    n := 1
    for n < 100 {
        n *= 2
    }
}
```

`for {}` : 무한 루프 

`for 인덱스, 요소값 := range 컬렉션 {}`

```go
names := []string{"홍길동", "이순신", "강감찬"}

for index, name := range names {
    println(index, name)
}
```

`break continue goto`
```go
package main 
func main() {
    var a = 1
    for a < 15 {
        if a == 5 {
            a += a
            continue     // for문의 나머지 사항들을 다 건너뛰고 다음 반복을 시행
        }
        a++
        
        if a > 10 {
            break        // for문 빠져나옴
        }
    }
    if a == 11 {  
        goto END         // goto가 지정하는 문장으로 이동 가능, indent가 중요
    }
    println(a)

END:
    println("End")
}
```

`break label` : 보통 break label은 직속 for문을 가르킨다. 결과적으로 label 다음 문장을 수행
```go
package main

func main() {
    i := 0

L1:
    for {
        if i == 0 {
            break L1
        }
    }
}
```