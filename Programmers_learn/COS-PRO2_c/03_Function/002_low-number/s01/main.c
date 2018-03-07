#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int solution(int arr[][4], int n, int k) {
    int *box = (int *)malloc(sizeof(int)*(n*4));

    int answer = 0, i, j, b = 0, tmp;

    for(i = 0; i < n; ++i) {
        for(j = 0; j < 4; ++j) {
            box[b] = arr[i][j];
            b++;
        }
    }   

    for(i = 0; i < b-1; ++i) {
        for(j = i+1; j < b; ++j) {
            if(box[i] > box[j]) {
                tmp = box[i];
                box[i] = box[j];
                box[j] = tmp;
            }
        }
    }   
    answer = box[k-1];
    return answer;
}

int main() {
	int arr[][4] = {
		{5, 12, 4, 31},
		{24, 13, 11, 2},
		{43, 44, 19, 26},
		{33, 65, 20, 21}
	}; 
	int n = 4;
	int k = 4;
	int ex = solution(arr, n, k);	// 11
	printf("%d", ex);
}
