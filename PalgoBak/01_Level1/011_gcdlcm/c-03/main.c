#include <stdio.h>

typedef struct GcdLcm
{
	int gcd;
	int lcm;
} gl;

int gcd(int n1, int n2);
int lcm(int n1, int n2);
gl gcdlcm(int n1, int n2);

int main() {
	gl ex1 = gcdlcm(3, 12);	// 3, 12
	printf("%d, %d", ex1.gcd, ex1.lcm);

	gl ex2 = gcdlcm(4, 7);	// 1, 28
	printf("\n%d, %d", ex2.gcd, ex2.lcm);
	
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
	max = (n1 > n2) ? n1 : n2;

	do {
		if(max%n1 == 0 && max%n2 == 0) { break; }
		else { ++max; }
	}
	while(1);

	return max;
}

gl gcdlcm(int n1, int n2) {
	gl box;
	box.gcd = gcd(n1, n2);
	box.lcm = lcm(n1, n2);
	
	return box;
}
