package main

func main(){
	map1()
	map2()
	map3()
	map4()
}

// map 선언 방법
func map1(){
	// 1
	var idMap map[int]string
	idMap = make(map[int]string)

	//2
	idMap2 := map[string]string{
		"GOO": "GOOGLE",
		"MS": "MICROSOFT",
		"FB": "FACEBOOK",
	}

	println(idMap, idMap2)
}

func map2(){
	var m map[int]string
 
    m = make(map[int]string)
    //추가 혹은 갱신
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"
 
    // 키에 대한 값 읽기
    str := m[134]
    println(str)
 
    noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
    println(noData)
 
    // 삭제
    delete(m, 777)
}

func map3(){
	tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",
    }
 
    // map 키 체크
    val, exists := tickers["MSFT"]

    if !exists {
        println("No MSFT ticker")
    }

	println(val)
}

func map4(){
	myMap := map[string]string{
        "A": "Apple",
        "B": "Banana",
        "C": "Charlie",
    }
 
    // for range 문을 사용하여 모든 맵 요소 출력
    // Map은 unordered 이므로 순서는 무작위
    for key, val := range myMap {
        println(key, val)
    }
}