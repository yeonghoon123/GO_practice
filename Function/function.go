package main

func main() {
    msg := "Hello"
    say(msg)
	say2(&msg)
	say3("This", "is", "a", "book")
	total := sum(1, 7, 3, 5, 9)
	println(total)
	count, total2 := sum2(1, 7, 3, 5, 9)
	count2, total3 :=sum3(1, 7, 3, 5, 9)
    println(count, total2)
	println(count2, total3)
}
 
func say(msg string) {
    println(msg)
}

func say2(msg *string) {
    println(*msg)
    *msg = "Changed" //메시지 변경
}

func say3(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}
 
func sum(nums ...int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    return s
}

func sum2(nums ...int) (int, int) {
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}

func sum3(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}