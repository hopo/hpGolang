#include <stdio.h>

void sumMatrix(int a[2][2], int b[2][2], int sum[2][2]);
void display(int marr[2][2]);

int main() {
	int a[2][2] = {{1, 2}, {2, 3}};
	int b[2][2] = {{3, 4}, {5, 6}};
	int sum[2][2];
	sumMatrix(a, b, sum); // {{4, 6}, {{7, 9}}
	display(sum);

	return 0;
}

void sumMatrix(int a[2][2], int b[2][2], int sum[2][2]) {
	int i, j, lnth;
	lnth = 2;
	for(i = 0; i < lnth; i++)
		for(j = 0; j < lnth; j++) {
			sum[i][j] = a[i][j]+b[i][j];
		}
}

void display(int marr[2][2]) {
	int i, j;
	printf("{");
	for(i = 0; i < 2; i++) {
		printf("{");
		for(j = 0; j < 2; j++) {
			printf(" %d ", marr[i][j]);
		}
		printf("}");
	}
	printf("}");
}
