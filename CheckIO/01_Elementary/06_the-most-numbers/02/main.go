package main

import "fmt"

func mostNumbers(nums ...float64) float64 {
	if nums != nil {
		lar, sma := nums[0], nums[0]
		for i := 1; i < len(nums); i++ {
			v := nums[i]
			switch {
			case v > lar:
				lar = v
			case v < sma:
				sma = v
			}
		}
		return lar - sma
	}
	return 0
}

func main() {
	fmt.Println(mostNumbers(1, 2, 3))                 //2, "3-1=2"
	fmt.Println(mostNumbers(5, -5))                   //10, "5-(-5)=10"
	fmt.Println(mostNumbers(10.2, -2.2, 0, 1.1, 0.5)) //12.4 "10.2-(-2.2)=12.4"
	fmt.Println(mostNumbers())                        //0 "Empty"
}
