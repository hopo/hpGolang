#include <stdio.h>


void reverseInt(int);

int main() {
	int num = 118372;
	reverseInt(num);	// 873211
}

void reverseInt(int num) {
	int dv = 10, box[12];

	while(1) {
		if(dv-num > 0) {
			dv /= 10;
			break;
		}
		dv *= 10;
	}

	printf("%d ", num/dv);
}


/*
func reverseInt(n int) int {
	isl := split_int_each(n)
	sorted := sortRv_isl(isl)
	ret := each_int_makeNum(sorted)

	return ret
}

func split_int_each(n int) []int {
	numlen := 10
	for {
		if n-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}

	var ret []int
	for {
		x := n / numlen
		ret = append(ret, x)
		r := n % numlen
		n = r
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}

func sortRv_isl(isl []int) []int {
	for i, _ := range isl {
		for j, _ := range isl {
			if i != j && isl[i] > isl[j] {
				isl[i], isl[j] = isl[j], isl[i]
			}
		}
	}
	return isl
}

func each_int_makeNum(isl []int) int {
	l := len(isl)
	numlen := pow(10, l-1)
	var ret int
	for _, v := range isl {
		ret += v * numlen
		numlen /= 10
	}
	return ret
}

func pow(n, p int) int {
	ret := 1
	for i := 0; i < p; i++ {
		ret *= n
	}
	return ret
}
*/
