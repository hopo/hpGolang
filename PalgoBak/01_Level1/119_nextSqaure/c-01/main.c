#include <stdio.h>
#include <math.h>

int nextSqaure(int);

int main() {
	int d1 = 121; 
	int ex1 = nextSqaure(d1);	// 144, 12*12
	printf("%d", ex1);

	int d2 = 4; 
	int ex2 = nextSqaure(d2);	// 9, 3*3
	printf("\n%d", ex2);

	int d3 = 5;
	int ex3 = nextSqaure(d3);	// 0, "no"
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
