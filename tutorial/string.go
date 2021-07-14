package main

import "fmt"

func main() {

	rawLiteral := `아리랑\n아이랑\n아라리요`

	interLiteral := "아리랑아리랑\n" + "아라리요" // + 연산자를 사용하면 2줄에 걸쳐 사용 가능

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}
