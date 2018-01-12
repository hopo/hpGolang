#include <stdio.h>

#define ROWS 2
#define COLS 2

int *sumMatrix(int a[][COLS], int b[][COLS]) {
	static int ret[ROWS][COLS];

	int i, j;
	for(i = 0; i < COLS; i++) {
		for(j = 0; j < COLS; j++) {
			ret[i][j] = a[i][j]+b[i][j];
		}
	}

	return ret;
}

int main() {
	int a[][COLS] = {{1, 2}, {2, 3}};
	int b[][COLS] = {{3, 4}, {5, 6}};

	int (*ret)[COLS];
	ret = sumMatrix(a, b); // {{4, 6}, {{7, 9}}

	int i, j;
	for(i = 0; i < COLS; i++) {
		for(j = 0; j < COLS; j++) {
			printf("%d ", ret[i][j]);
		}
	}

	return 0;
}
