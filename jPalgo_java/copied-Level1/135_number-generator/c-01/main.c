#include <stdio.h>

int *number_generator(int, int);

int main() {
	int *ex1 = number_generator(3, 5); // [3 6 9 12 15]
	printf("%d %d %d %d %d", ex1[0], ex1[1], ex1[2], ex1[3], ex1[4]);

	int *ex2 = number_generator(4, 3); // [4 8 12]
	printf("\n%d %d %d", ex2[0], ex2[1], ex2[2]);

	return 0;
}

int *number_generator(int x, int n) {
	int i;
	int ret[n];
	for(i = 0; i < n; i++) {
		ret[i] = x*(i+1);	
	}

	int *dt;
	dt = ret;

	return dt;
}
