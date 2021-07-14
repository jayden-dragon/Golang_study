package main // 1. package 정의

import "fmt" // 2. package import

/*
func main(){  // 3. 함수 정의, 선언
  fmt.Println("Hello World")
}
*/

/*
func main() {
	var result int
	result = 1 + 2
	fmt.Println(result)
}
*/

func main(){

  rawLiteral := '아리랑\n아리랑\n아라리요.'

  interLiteral := "아리랑아리랑\n" + "아라리요"

  fmt.Println(rawLiteral)
  fmt.Println()
  fmt.Println(interLiteral)

// + 연산자를 사용하면 2줄에 걸쳐 사용 가능
/* 출력
 아리랑\n
 아리랑\n
 아리리요
 */
}
