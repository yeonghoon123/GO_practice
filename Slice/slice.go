package main
import "fmt"
 
func main() {
	slice1()
}

// 기본 슬라이스
func slice1() {
	// slice 선언 방법 1
	var a []int        //슬라이스 변수 선언
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	a[1] = 10
	fmt.Println(a)     // [1, 10, 3]출력

	// slice 선언 방법 2
	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10

	var s2 []int

	if s2 == nil {
		println("Nil Slice")
	}

	println(len(s2), cap(s2)) // 모두 0
}

// 부분 슬라이스
func slice2(){
	s := []int{0,1,2,3,4,5}
	s = s[2:5]
	s = s[2:5]     // 2, 3, 4
	s = s[1:]      // 3, 4
	fmt.Println(s)
}

// 슬라이스 추가
func slice3(){
	// 슬라이스 추가
	s := []int{0, 1}
	
	s = append(s, 2)

	s = append(s, 3, 4, 5)

	fmt.Println(s)

	// 슬라이스 추가시 배열길이와, 용량 비교
	// len=0, cap=3 인 슬라이스
    sliceA := make([]int, 0, 3)
 
    // 계속 한 요소씩 추가
    for i := 1; i <= 15; i++ {
        sliceA = append(sliceA, i)
        // 슬라이스 길이와 용량 확인
        fmt.Println(len(sliceA), cap(sliceA))
    }
 
    fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력 
}

// 슬라이스 병합
func slice4(){
	sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
 
    sliceA = append(sliceA, sliceB...)
    //sliceA = append(sliceA, 4, 5, 6)
 
    fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
}

// 슬라이스 복사
func slice5(){
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source) *2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target)) // 3, 6 출력
}

