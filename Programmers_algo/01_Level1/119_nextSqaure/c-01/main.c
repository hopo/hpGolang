#include <stdio.h>
#include <math.h>

int nextSqaure(int);

int main() {
	int ex1 = nextSqaure(121); // 144, 12*12
	printf("%d", ex1);

	int ex2 = nextSqaure(4);   // 9, 3*3
	printf("\n%d", ex2);

	int ex3 = nextSqaure(5);   // 0, "no"
	printf("\n%d", ex3);

	return 0;
}

int nextSqaure(int n) {
	float sr;
	sr = sqrt(n);
	if(sr != (int)sr) { return 0; }

	int srp, ret;
	srp = sr + 1;
	ret = pow(srp, 2);

	return ret;
}
