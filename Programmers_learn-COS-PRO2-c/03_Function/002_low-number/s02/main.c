#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int solution(int arr[][4], int n, int k) {
    int i, j, tmp, bi = 0, answer = 0, linear[n*4];

    for(i = 0; i < n; i++)
        for(j = 0; j < 4; j++)
            linear[bi++] = arr[i][j];

    for(i = 0; i < bi-1; i++) 
        for(j = i+1; j < bi; j++)
            if(linear[i] > linear[j]) {
                tmp = linear[i];
                linear[i] = linear[j];
                linear[j] = tmp;
            }
       
    answer = linear[k-1];

    return answer;
}

int main() {
	int ex, n = 4, k = 4;
	int arr[][4] = {
		{5, 12, 4, 31},
		{24, 13, 11, 2},
		{43, 44, 19, 26},
		{33, 65, 20, 21}
	}; 
	ex = solution(arr, n, k);	// 11
	printf("%d", ex);
}
