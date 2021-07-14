# Array
`var  변수명 [배열크기]데이터타입`

```go
package main

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}    // 배열 크기 자동 설정 

	var multiArray [3][4][5]int   // 다차원 배열 선언 
	multiArray[0][1][2] = 10

	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},     // 끝에 콤마 추가해야댐
	}
	println(a[1][2])
}

```


