#include <stdio.h>

void digit_reverse(int num, int *ret);

int main() {
	int num1 = 12345, ret1[16];
	digit_reverse(num1, ret1); // [5 4 3 2 1]
	printf("[%d %d %d %d %d]\n", ret1[0], ret1[1], ret1[2], ret1[3], ret1[4]);

	int num2 = 321, ret2[16];
	digit_reverse(num2, ret2);   // [1 2 3]
	printf("[%d %d %d]", ret2[0], ret2[1], ret2[2]);

	return 0;
}

void digit_reverse(int num, int *ret) {
	int ea, nl = 10, lnth = 1;
	while(1) {
		if(nl-num > 0) {
			nl /= 10;
			break;
		}
		nl *= 10;
		lnth++;
	}

	while(1) {
		ea = num / nl;
		ret[lnth-1] = ea;
		lnth--;
		num %= nl;
		nl /= 10;
		if(nl < 1) { break; }
	}
}
