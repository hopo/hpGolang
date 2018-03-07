#include <stdio.h>

void water_melon(int n, char ret[]);

int main() {
	int n1 = 3;
	char ret1[n1+1];
	water_melon(n1, ret1);	// "WMW"
	printf("%s", ret1);

	int n2 = 4;
	char ret2[n2+1];
	water_melon(n2, ret2);	// "WMWM"
	printf("\n%s", ret2);

	return 0;
}

void water_melon(int n, char ret[]) {
	int r = 0;
	for(int i = 1; i < n+1; ++i) {
		(i%2 == 1)? (ret[r] = 'W') : (ret[r] = 'M');
		r++;
	}
	ret[r] = '\0';
}
