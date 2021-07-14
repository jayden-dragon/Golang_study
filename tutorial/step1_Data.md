# 1. What is Go?

-	Go : 컴파일 언어 - Go Toolchain을 이용해서 소스 파일과 다양한 dependency를 컴퓨터의 기본 명령어로 변환
-	go run or go run <file> : 한 개 이상의 소스파일이나 해당 소스파일을 컴파일한고 실행
-	go build <file> : 해당 파일의 exe파일 생성

Hello World
-----------

```go
package main

import "fmt" // 정형 데이터 입력 및 출력을 위한 함수 library

func main() {
	fmt.Prinln("Hello World!")
}
```

Syntax, Formatting
------------------

-	문장이 한 줄에 두개 이상 나오는 경우 외에는 문장이나 선언의 끝에 세미콜론 x

# 2. Variable

```go

var a int                   // 변수 선언 시 var 변수명 변수타입

var f float32 = 11.         // 선언과 동시에 초기값 설정 가능

a = 10                      // 선언 후 해당 타입의 값 할당 가능
f = 12.0

var i, j, k int            // 변수 여러개를 동시에 선언 가능

var i, j, k int = 1, 2, 3  // 동시 선언과 동시에 초기값 동시 설정 가능


// 변수 선언 시 초기값을 지정하지 않으면 Go는 zero value를 default로 할당
// bool에서는 false, string에서는 ""(빈문자열)

var i = 1 var s = "hi"     // Go에서는 할당되는 값을 보고 데이터 타입을 추론
```

Short Assignment Statement (:=)
-------------------------------

-	func 내에서 변수를 사용하는 간단한 방식`go i := 1      // var i = 1 대신 사용 가능`

# 3. Constant

```go
const c int = 10
const s string = "HI"

const c = 10
const s = "HI"

const ( // 여러 개의 상수들 묶어서 지정 가능
	Visa   = "Visa"
	Master = "Master Card"
	Amex   = "American Express"
)

const ( // iota라는 identifier를 사용하면 0부터 순차적인 값 부여
	Apple  = iota // 0
	Grape         // 1
	Banana        // 2
)
```

# 4. Data Type

-	bool
-	string
-	int int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64
-	float & complex float32 float64 complex64 complex128
-	etc byte : uint8와 동일하며 byte 코드에 사용 rune : int32와 동일하며 유니코드 코드포인트에 사용

String
------

-	*Raw String literal('')* '' 내부에 있는 문자열은 별도의 해석 없이 raw string 그대로 인식 ex) 문자열 내부에 \n이 있어도 new line으로 해석 x

-	*Interpreted String Literal("")* 복수의 라인에 걸쳐 사용 불가 "" 내부의 escape 문자열들은 특별한 의미로 해석 ex) \n는 new line으로 해석, + 연산자로 결합 사용 가능

```go
package main

import "fmt"

func main(){

  rawLiteral := '아리랑\n
                아리랑\n
                아라리요.'

  interLiteral := "아리랑아리랑\n" + "아라리요"   // + 연산자를 사용하면 2줄에 걸쳐 사용 가능

  fmt.Println(rawLiteral)
  fmt.Println()
  fmt.Println(interLiteral)


/* 출력   
아리랑\n아이랑\n아라리요

아리랑아리랑
아라리요
 */
}
```

Data Type conversion
--------------------

T(v) : Type conversion - T : 변환하고자 하는 데이터 타입 - v : 변환될 값 ex) float32(100), []byte("ABC") : 문자열을 byte배열로

```go

func main() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
```
