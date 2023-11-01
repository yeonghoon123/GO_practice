package main

import(
	"fmt"
)

func main(){
	// Back Quoute 사용
var test string =
`테스트\n테스트
테스트
`

// 이중인용부호 사용
var test2 string = "테스트\n테스트테스트"

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


fmt.Println(test, test2)
fmt.Println(i,y,z,str,bytes,str2)
}