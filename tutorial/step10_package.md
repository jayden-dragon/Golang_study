# Package
- Go는 package를 통해 코드의 모듈화, 코ㅗ드의 재사용 기능을 제공한다. Go는 package를 사용해서 작은 단위의 component를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장 
- Go는 실제 프로그램 개발에 필요한 많은 package들을 표준 library로 제공한다. 이러한 패키지들은 GOROOT/pkg 안에 존재한다. GOROOT환경변수는 Go 설치 디렉토리를 가리ㅣ키는데, 보통 Go 설치시 자동으로 추가된다.

- https://golang.org/pkg : Go 패키지 자세한 설명

## Package scope 
- package 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재하는데, 이들의 이름의 첫문자를 대문자로 시작하면 public으로 사용 가능하고 소문자일 경우 non-public으로 패키지 내부에서만 사용 가능하다.

## Package init & alias
init() : 패키지 실행 시 처음으로 호출되기에 따로 함수 호출 없이 자동으로 호출
```go
package testlib

var pop map[string]string

func init() {
    pop = make(map[string]string)
}
```
경우에 따라, package를 import 하면서 package 내부에 init() 함수만을 호출하고자 하는 경우가 있다. 이때 import 시 _라는 alias를 지정한다.


## User define package 
- 개발자는 사용자 정의 패키지를 만들어 재사용 가능한 component를 만들어 사용할 수 있다.
- 사용자 정의 library 패키지는 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성한다. 
- 하나의 서브 폴더 안에 있는 .go파일들은 동일한 패키지명을 가지며, 패키지명은 해당 폴더의 이름과 같게 한다. 즉, 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶인다.

src 폴더 내부의 testlib라는 폴더를 만들고 안에 music.go라는 파일을 저장 -> import testlib으로 패키지를 import 한다.

```go
// music.go
package testlib

import "fmt"

var pop map[string]string

func init() {
    pop = make(map[string]string)
    pop["Adele"] = "Hello"
    pop["Alicia Keys"] = "Fallin"
    pop["john Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
    return pop[singer]
}

func getKeys() {
    for _, kv := range pop {
        fmt.Println(kv)
    }
}
```
*Optional* <br/>
사이즈가 큰 복잡한 library 같은 경우, `go install` 명령을 사용하여 library를 컴파일하여 cache할 수 있는데, 이럴 경우 다음 빌드 시 빌드타임을 크게 줄일 수 있다. Go 패키지를 빌드하고 /pkg 폴더에 인스톨하기 위해서 cd .\testlib\ 하고 go install 하면 pkg 내부에 testlib.a라는 파일이 생성된다. 

```go
package main

import "testlib"

func main() {
    song := testlib.GetMusic("Alicia Keys")
    println(song)
}
```