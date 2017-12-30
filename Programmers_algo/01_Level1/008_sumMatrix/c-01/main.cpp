// ing
// https://programmers.co.kr/learn/challenge_codes/148

#include <stdio.h>

int sumMatrix(int, int);

int main() {
	int a[2][2] = {{1, 2}, {2, 3}};
	int b[2][2] = {{3, 4}, {5, 6}};
	int ex1 = *sumMatrix(&a, &b); // {{4, 6}, {{7, 9}}

	for(int i = 0; i < answer.size(); i++) {
		for(int j = 0; j < answer[0].size(); j++) {
			printf(answer[i][j]);
		}
		printf("\n");
	}
}

int *sumMatrix(int *A[][2], int *B[][2]) {
	int answer[2][2] = *A;
	int st = A.size();

	int i, j;

	for(i = 0; i < st; i++) {
		for(j = 0; j < st; j++) {
			answer[i][j] = A[i][j]+B[i][j];
		}
	}	

	return &answer;
}
