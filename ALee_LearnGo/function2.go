package main

func main() {
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)
}

func sum(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}
