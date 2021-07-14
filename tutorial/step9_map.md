# Map 
- Map은 key에 대응하는 value를 신속히 찾는 hash table을 구현한 자료구조 <br/>
- `var name map[key타입]value타입` 형태로 선언
``` 
var idMap map[int]string
```
- 이때 선언된 변수 idMap은 reference 타입이며 초기 nil 값을 가진다. -> Nil map
- map을 초기화하기 위해선 make()를 사용
```go
idMap = make(map[int]string)
```
- literal 초기화도 가능 : `map[key타입]value타입 {key:value}`
```go
tickers := map[string]string {
    "GOOG" : "Google Inc",
    "MSFT" : "Microsoft",
    "FB" : "FaceBook",
}
```

- make() 초기화 후 아무 데이터도 없는 상태, 새로운 데이터 추가를 위해 `map변수[key] = value` 와 같이 해당 키에 값을 할당하면 된다.

```go
package main

func main() {
    var m map[int]string

    m = make(map[int]string)
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"

    str := m[134]
    println(str)

    noData := m[999]   // 비어있으면 nil or zero
    println(noData)

    delete(m, 777)   // 삭제 
}
```

## Map key check

```go
package main

func main() {
    tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",
    }

    // map 키 체크
    val, exists := tickers["MSFT"]  // 2개의 return값이 존재. key & key 유무 bool
    if !exists {
        println("No MSFT Ticker")
    } else {
        println(val, exists)
    }
}
```

# for loop를 사용한 map 열거
```go
package main

import "fmt"

func main() {
    myMap := map[string]string {
        "A" : "Apple",
        "B" : "Banana",
        "C" : "Charlie",
    }

    for key, val := range myMap {
        fmt.Println(key, val)
    }
}
```