package main

func main() {
 names := []string{"홍길동", "이순신", "강감찬"} //이름 배열

 // 기본 반복문
 sum := 0
        for i := 1; i <= 100; i++ {
            sum += i
        }
        println(sum)

// 조건식만 있는 반복문
 n := 1
        for n < 100 {
            n *= 2
            //if n > 90 {
            //   break
            //}
        }
        println(n)

// 무한루프
//  for {
// 	println("Infinite loop")
// }

// for range문
    for index, name := range names {
        println(index, name)
    }

// break, continue, goto문
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break  //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
    }