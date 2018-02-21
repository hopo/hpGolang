package main

/*
int four() {
return 4;
}
*/
import "C"
import "fmt"

func main() {
	f := C.four()
	fmt.Println(f)
}
