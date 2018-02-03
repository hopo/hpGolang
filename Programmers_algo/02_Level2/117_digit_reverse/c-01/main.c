// ing
#include <stdio.h>

void digit_reverse(int num, int *ret);

int main() {
	int num1 = 12345, ret1[16];
	digit_reverse(num1, ret1); // [5 4 3 2 1]
	printf("[%d %d %d %d %d]", ret1[0], ret1[1], ret1[2], ret1[3], ret1[4]);

	/*
	int num2 = 321, ret2[16];
	digit_reverse(num2, ret2);   // [1 2 3]
	*/

	return 0;
}

void digit_reverse(int num, int*ret) {
	
}

/*

func digit_reverse(n int) []int {
	isl := int_convert_each(n)
	var ret []int
	for i := 0; i < len(isl); i++ {
		j := len(isl) - 1 - i
		ret = append(ret, isl[j])
	}
	return ret

}

func int_convert_each(num int) []int {
	numlen := 10
	for {
		if num-numlen < 0 {
			numlen /= 10
			break
		}
		numlen *= 10
	}
	var ret []int
	var ea int
	for {
		ea = num / numlen
		ret = append(ret, ea)
		num %= numlen
		numlen /= 10
		if numlen < 1 {
			break
		}
	}
	return ret
}
*/
