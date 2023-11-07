package main

import(
	"fmt"
)

type person struct {
	name string
	age int
}

func newPerson() *person {
	p := person{}

	return &p
}

// person객체 생성 방법
func main(){
	// 방법1
	p := person{}

	//필드값 설정
	p.name = "kyh"
	p.age = 23


	// 방법2
	var p1 person
	p1 = person{"kwb", 43}
	p2 := person{"jws", 13}

	// 방법3 
	p3 := new(person)
	p3.name = "jbr"

	// 방법4 (생성자 함수)
	p4 := newPerson()
	p4.name = "wow"
	p4.age = 31

	fmt.Println(p, p1, p2, p3, p4)
}