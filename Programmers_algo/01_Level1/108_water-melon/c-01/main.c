#include <stdio.h>

char *water_melon(int);

int main() {
	char *ex1 = water_melon(3);	// "WMW"
	printf("%s", ex1);

	char *ex2 = water_melon(4);	// "WMWM"
	printf("\n%s", ex2);

	return 0;
}

char *water_melon(int num) {
	char ret[num+1];
	int r = 0;
	for(int i = 1; i < num+1; ++i) {
		(i%2 == 1)? (ret[r] = 'W') : (ret[r] = 'M');
		r++;
	}
	ret[r] = '\0';

	char *dt = ret;

	return dt;
}
