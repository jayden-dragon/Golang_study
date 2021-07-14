# 1.Operator

## 산술연산자

-	사칙연산 : = - * / %
-	증감연산 : ++ --

## 관계연산자

-	a == b
-	a != b
-	a >= b

## 논리연산자

-	AND, OR, NOT
-	a && b
-	a || !(c && b)

## Bitwise 연산자

-	bit 단위 연산에 사용
-	AND, OR, XOR, binary shift
-	c = (a & b) \<< 5

## 할당연산자

-	할당 연산자는 = 오ㅚ에도 사칙연산 및 비트연산을 축약한 연산자
-	a = 100
-	a *= 10
-	a >>= 2
-	a |= 1

## 포인터연산자

```go
var k int = 10
var p = &k   // k의 주소를 할당
printlnl(*p) // p가 가리키는 주소에 있는 실제 내용 출력
```
