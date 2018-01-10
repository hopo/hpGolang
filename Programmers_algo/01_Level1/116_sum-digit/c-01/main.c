#include <stdio.h>

int sum_digit(int);

int main() {
	int ex1 = sum_digit(123);       // 6
	printf("%d", ex1);
	int ex2 = sum_digit(999999999); // 81
	printf("\n%d", ex2);

	return 0;
}

int sum_digit(int n) {
	int ret, z, di;
	ret = z = 0;
	di = 1;

	while(1) {
		if(n-di < 0) {
			di = di/10;
			break;
		}
		di *= 10;
	}

	while(1) {
		if(di == 1) {
			ret += n;
			break;
		}
		z = n/di;
		ret += z;
		n = n%di;
		di = di/10;
	}

	return ret;
}
