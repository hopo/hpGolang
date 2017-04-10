package main

func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)
}

func sum(nums ...int) (int, int) {
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}
