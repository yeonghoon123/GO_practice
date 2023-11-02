package main

import(
	"fmt"
)

func main(){
	var testVal1 = 1

	// if문
	if testVal1 != 1{
		fmt.Println("success")
	} else if testVal1 == 1{
		fmt.Println("testVal1 = 1")
	} else {
		fmt.Println("fail")
	}

	if val := testVal1 * 2; val < 3 {
		println(val)
	}

	// switch문
	switch testVal1 {
	case 1: fmt.Println(testVal1)
	case 2: fmt.Println("2 log")
	}

	// 인자값 없음
	switch {
    case testVal1 >= 90:
        println("A")
    case testVal1 >= 80:
        println("B")
    case testVal1 >= 70:
        println("C")
    case testVal1 >= 60:
        println("D")
    default:
        println("No Hope")
    }

	swtichCase1(21)
	check(2)
}

// 변수 타입형으로 구분 가능  
func swtichCase1(value interface{}){
	switch value.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}
}

// break값 없이 출력
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}