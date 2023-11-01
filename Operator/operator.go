package main

import(
	"fmt"
)

func main(){
	var k int = 10
	var p = &k
	fmt.Println(p) // p가 가리키는 주소에 있는 실제 내용을 출력
}