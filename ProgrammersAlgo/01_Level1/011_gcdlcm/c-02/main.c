#include <stdio.h>

int gcd(int n1, int n2);
int lcm(int n1, int n2);
int* gcdlcm(int n1, int n2);

int main() {
	int* ex1 = gcdlcm(3, 12);	// 3, 12
	printf("%d, %d", *ex1, *(ex1+1));

	int *ex2 = gcdlcm(4, 7);	// 1, 28
	printf("\n%d, %d", *ex2, *(ex2+1));
	
	return 0;
}

int gcd(int n1, int n2) {
	// bigger - smaller ?
	while(n1 != n2) {
		if(n1 > n2) { n1 -= n2; }
		else { n2 -= n1; }
	} 

	return n1;
}

int lcm(int n1, int n2) {
	int max;

	// maximum value between n1 and n2 is stored in max
	max = (n1 > n2)? n1 : n2;

	do {
		if(max%n1 == 0 && max%n2 == 0) { break; }
		else { ++max; }
	}
	while(1);

	return max;
}

int* gcdlcm(int n1, int n2) {
	int r1, r2;
	r1 = gcd(n1, n2);
	r2 = lcm(n1, n2);

	int ret[] = {r1, r2};

	int* dt = ret;

	return dt;
}

