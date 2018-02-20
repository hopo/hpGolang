#include <stdio.h>

void productMatrix(int mata[2][2], int matb[2][2], int ret[2][2]);

int main() {
	int mata[2][2] = {{1, 2}, {2, 3}};
	int matb[2][2] = {{3, 4}, {5, 6}};
	int ret[2][2];
	productMatrix(mata, matb, ret);	// {{13, 16}, {21, 26}}
	printf("{{%d, %d}, {%d, %d}}", ret[0][0], ret[0][1], ret[1][0], ret[1][1]);

	return 0;
}

void productMatrix(int mata[2][2], int matb[2][2], int ret[2][2]) {
	int i, j, lnth = 2;
	for(i = 0; i < lnth; i++) {
		for(j = 0; j < lnth; j++) {
			ret[i][j] = (mata[i][0]*matb[0][j]) + (mata[i][1]*matb[1][j]);
		}
	}
}
