# Slice
- 배열은 고정된 크기 안에 동일한 타입의 data를 연속적으로 저장하지만, 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능 x <br/>
- Slice는 내부적으로 배열에 기초하여 만들어 졌지만 배열의 이런 제약점들을 넘어 개발자에게 편리하고 유용한 기능들을 제공. slice는 배열과 달리 고정된 크기를 미리 지정 x, 차후 크기를 동적으로 변경 가능, 또한 부분 배열 발췌도 가능

- `var v []T` : v는 변수명 T는 데이터 타입

```go
package main
import "fmt"

func main(){
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)
}
```
- Go의 내장함수 make()를 통해서 slice 생성 가능
- make()를 통해 생성하면 slice의 길이와 용량을 임의로 지정할 수 있는 장점
- `a := make([]int, 5, 10)` : 슬라이스 타입, 슬라이스 길이, 슬라이스 용량(내부 배열의 최대 길이) 만약 용량값을 생략하면 길이값과 동일한 값이 적용된다.
- len(), cap()을 써서 확인 가능

```go
func main(){
	s := make([]int, 5, 10)
	println(len(s), cap(s))
}


// 슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0인 슬라이스 생성하는데 이를 Nil Slice라고 한다
func main() {
	var s []int

	if s == nil {
		println("Nil Slice")
	}
	println(len(s), cap(s))
}

```

# Sub-Slice
Slice에서 일부를 발췌하여 부분 슬라이스를 만들 수 있다. "슬라이스[처음인덱스:마지막인덱스]" 형식으로 만든다. 
```go
package main

import "fmt"

func main() {
    s := []int{0, 1, 2, 3, 4, 5}
    s = s[2:5]
    fmt.Println(s) 
    // output : 2, 3, 4
    
    s = s[1:]
    fmt.Println(s)
    //output : 3, 4
}
```

# Slice append & copy
배열은 고정된 크기로 인해 데이터를 임의로 추가 불가 but 슬라이스는 자유롭게 새로운 요소를 추가 가능. 이때 append() 내장함수를 통해 가능 `append(s, 2)` : s: 슬라이스 객체, 2:추가할 

```go
func main() {
    s := []int{0, 1}

    s = append(s, 2)

    s = append(s, 3, 4, 5)
    
    fmt.Println(s)
    //output : 0, 1, 2, 3, 4, 5
}

// 만약 slice 용량이 부족한데 append할 경우 결과적으로 새로운 slice를 할당

package main

import "fmt"

func main() {
    sliceA := make([]int, 0, 3)

    for i := 1; i <= 15; i++ {
        sliceA = append(slliceA, i)
        fmt.Println(len(sliceA), cap(sliceA))
    }

    fmt.Println(sliceA)
}
```
```go
package main
import "fmt"

func main(){
    sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}

    sliceA = append(sliceA, sliceB...)
    //sliceA = append(sliceA, 4, 5, 6)
    fmt.Println(sliceA)   // output 1, 2, 3,4 ,5 ,6 
}


func main() {
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    
    fmt.Println(target)  // 0 ,1 ,2 출력
    println(len(target), cap(target))  // 3, 6 출력
}
```
