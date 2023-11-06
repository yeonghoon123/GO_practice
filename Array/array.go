package main

func main() {
    var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 1
    a[1] = 2
    a[2] = 3
    println(a[1]) // 2 출력

    // 배열에 초기화
    var a1 = [3]int{1, 2, 3}
    var a2 = [...]int{1, 2, 3} //배열크기 자동으로

	println(a1[1], a2[2])

    // 다차원 배열
    var multiArray [3][4][5]int  // 정의
    multiArray[0][1][2] = 10     // 사용

    // 다차원 배열의 초기화
    var a3 = [2][3]int{
            {1, 2, 3},
            {4, 5, 6},  //끝에 콤마 추가
        }

    println(a3[1][2])
}