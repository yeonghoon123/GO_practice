# GO lang practice

### 목차

1. [GO 공식 사이트](#go-공식-사이트)<br>
2. [GO란?](#go란)<br>
3. [GO 언어의 특성](#go-언어의-특성)<br>
4. [GO 기초 문법](#go-기초-문법)<br>
   4-1. [파일 생성 및 Hello World](#1-파일-생성-및-hello-world-출력)<br>
   4-2. [변수와 상수](#2-변수와-상수)<br>
   4-3. [데이터 타입](#3-go-데이터-타입)<br>
   4-4. [연산자](#4-연산자)<br>

<br>
<br>

### GO 공식 사이트

링크 : <a href='https://go.dev/'>https://go.dev</a>

<br>

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

### GO 기초 문법

---

### 1. 파일 생성 및 Hello World 출력

1.  <b>go 모듈 초기화</b>

    ```shell
    go mod init <모듈명>
    ```

    <br>

2.  \*.go 형식의 파일 생성 후 코드 기입</b><br>

    ```go
    package main

    import "fmt"

    func main(){
        fmt.Println("Hello World")
    }
    ```

    <br>

3.  명령어 입력

    ```shell
    go run .\*.go
    Hello World
    ```

<br>

### 2. 변수와 상수

1. <b>변수</b><br>

    ```go
    var a int = 10 // 정수타입
    var b float32 = 11.2 // 실수타입
    var i, j, k int = 1, 3, 2 // 같은 타입 복수 result) i = 1, j = 3, k = 2

    // func 내에서 사용가능한 변수
    func test(){
        q := 1 // var과 같은 Short Assignment Statement (:=) 'var q = 1' 와 같음
    }
    ```

<br>

2.  <b>상수</b>

    ```go
    const c int = 10 // 정수타입
    const s = "HI" // 문자열 타입

        // 상수를 묶어서 선언 가능
        const (
            korea = "Korea"
            japan = "Japan"
            usa = "USA"
        )

        /*
        상수값을 0 부터 순차적으로 부여할땐 iota라는 idntifier(식별자)를 사용가능,
        이 경우 iota가 지정된 Apple은 0, 나머지상수들은 순서대로 1씩 증가
        */

        const (
            Apple = iota // 0
            Grape        // 1
            Orange       // 2
        )
    ```

    <br>

3.  <b>예약 키워드</b><br>
    예약 키워드는 변수, 상수명으로 사용 불가<br>
    > break default func interface select case defer go map<br> struct chan else goto package switch const
    > fallthrough if range type continue for import return var

<br>

### 3. GO 데이터 타입

1.  <b>Boolean type</b><br>
    bool

    <br>

2.  <b>String type</b><br>
    string: string은 한번 생성되면 수정될 수 없는 Immutable(불변) type이다.
    <br>
    문자열 리터럴은 Back Quote(``), 이중인용부호("")를 사용한다.

    ```go
    // Back Quoute 사용
    var test string =
    `테스트\n테스트
    테스트
    `
    /* 출력:
    테스트\n테스트
    테스트
    */

    // 이중인용부호 사용
    var test2 string = "테스트\n테스트테스트"
    /*
    테스트
    테스트테스트
    */
    ```

    <br>

3.  <b>Integer type</b><br>
    int int8 int16 int32 int64<br>
    uint uint8 uint16 uint32 uint64 uintptr

    <br>

4.  <b>Float type</b><br>
    float32 float64 complex64 complex128

    <br>

5.  <b>Other type</b><br>
    byte: uint8과 동일하며 바이트 코드에 사용한다.<br>
    rune: int32과 동일하며 유니코드 코드포인트에 사용한다.

    <br>

    ```go
    // 데이터 타입 변환
    var i int = 100
    var y uint = uint(100)
    var z float32 = float32(100)

    /* 출력
    100, 100, 100
    */

    var str string = "ABC"
    var bytes = []byte(str)
    var str2 string = string(bytes)
    /*
    ABC [65 66 67] ABC
    */
    ```

<br>

### 4. 연산자

1.  <b>산술 연산자</b><br>
    산술 연산자는 사칙연산자(+, -, \*, /, %)와 증감연사자(++, --)를 사용

    ```go
        result = (a + b) / 5
        result++
    ```

    <br>

2.  <b>관계 연산자</b><br>
    관계 연산자는 서로의 크기를 비교, 동일함을 체크하는데 사용

    ```go
    a == b
    a != c
    a > = b
    ```

    <br>

3.  <b>논리 연산자</b><br>
    논리 연산자는 AND, OR, NOT을 표현하는데 사용

    ```go
    A && B
    A || !(C && B)
    ```

    <br>

4.  <b>Bitwise 연산자</b><br>
    Bitwise 연산자는 비트단위 연산을 위해 사용되는데, 바이너리 AND, OR, XOR와 바이너리 쉬프트 연산자가 있다.

    ```go
    c = (a & b) << 5
    ```

    <br>

5.  <b>할당 연산자</b><br>
    할당 연산자는 값을 할당하는 = 연산자 외에 사칙연산, 비트연산을 축약한 +=, &=, <<= 같은 연산자도 있다.

    ```go
    a = 100
    a *= 10
    a >>= 2
    a |= 1
    ```

    <br>

6.  <b>포인터 연산자</b><br>
    포인터 연산자는 C++과 같이 & 혹은 \* 을 사용하여 해당 변수의 주소를 얻어내거나, 이를 반대로 Dereference 할 때 사용한다.<br>
    GO는 포인터 연산자를 제공하지만 포인터 산술 즉 포인터에 더하고 빼는 기능은 제공하지 않음.
    ```go
    var k int = 10
    var p = &k
    println(*p) // p가 가리키는 주소에 있는 실제 내용을 출력
    ```

<br>

### 5. 조건문

1.  <b>if 문</b><br>
    if 문은 조건에 맞으면 {} 불럭안의 내용을 실행한다. GO의 if 조건문은 괄호로 둘러싸지 않아도 된다.<br>

    ```go
    if k == 1 {
    println("One")
    } else if k == 2 {  //같은 라인
    println("Two")
    } else {   //같은 라인
    println("Other")
    }
    ```

    <br>

2.  <b>switch문</b><br>
    여러 값을 비교 해야 하거나 다수의 조건식을 체크해야 하는 경우 switch문을 사용한다. switch문 뒤에 하나의 변수(Expression)을 지정하고, case문에 해당 변수가 가질 수 있는 값을 지정하여 조건식이 가능하다.<br>

    ```go
    var name string
    var category = 1

    switch category {
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
    println(name) // name = Paper Book
    ```

    | -------- | ------- |
    | switch 뒤에 expression이 없을수 있음|switch 키워드 뒤에 expression이 없어도 사용 가능하다. |
    |case문에 expression을 쓸수 있음|case 문에 복잡한 expression을 가질 수 있다.|
    | No default fall through | break를 사용하지 않아도 다음 case로 넘어가지 않는다. |
    | type switch | 변수의 type에 따라 case로 분기할 수 있다. |
