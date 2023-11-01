# GO lang practice

### GO란?

2007년 구글에서 개발한 언어이다.

<br>

### GO 언어의 특성

-   GO는 전통적인 컴파일, 링크 모델을 따르는 범용 프로그래밍 언어이다.
    <br>
-   일차적으로 시스템 프로그래밍을 위해 개발되었으며, C++,Java,Python의 장점들을 뽑아 만들어졌다.
    C++과 같이 컴파일러를 통해 컴파일 되며, 정접타입의 언어이고 또한 Java와 같이 Garbage Collection 기능을 제공한다.
    <br>
-   단순하고 간결한 프로그래밍 언어를 지향하였는데, Java의 절반에 해당하는 25개의 키워드만으로 프로그래밍이 가능하게 하였다.
    <br>
-   Communicating Sqeuential Processes(CSP) 스타일의 Concureent 프로그래밍을 지원한다.

<br>

### GO 공식 사이트

링크 : <a href='https://go.dev/'>https://go.dev</a>

<br>

### GO 기초 문법

---

### 1. 파일 생성 및 Hello World 출력

-   go 모듈 초기화

```shell
go mod init <모듈명>
```

<br>
-   \*.go 형식의 파일 생성 후 코드 기입

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello World")
}
```

<br>
- 명령어 입력

```shell
go run .\*.go
Hello World
```

<br>

### 2. 변수와 상수

```go
// 변수
var a int = 10 // 정수타입
var b float32 = 11.2 // 실수타입
var i, j, k int = 1, 3, 2 // 같은 타입 복수 result) i = 1, j = 3, k = 2

// func 내에서 사용가능한 변수
func test(){
    q := 1 // var과 같은 Short Assignment Statement (:=) 'var q = 1' 와 같음
}

// 상수
const c int = 10 // 정수타입
const s = "HI" // 문자열 타입

// 상수를 묶어서 선언 가능
const (
    korea = "Korea"
    japan = "Japan"
    usa = "USA"
)

/*
상수값을 0 부터 순차적으로 부여할땐 iota라는 idntifier를 사용가능,
이 경우 iota가 지정된 Apple은 0, 나머지상수들은 순서대로 1씩 증가
*/

const (
    Apple = iota // 0
    Grape        // 1
    Orange       // 2
)

/*
예약 키워드로 변수명, 상수명 사용 불가
break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var
*/
```
