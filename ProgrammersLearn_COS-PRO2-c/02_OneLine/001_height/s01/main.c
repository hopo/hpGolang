#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int solution(int height[], int height_len, int k) {
    int answer = 0;
    for(int i = 0; i < height_len; ++i)
        if(height[i] > k)
            answer++;
    return answer;
}

int main() {
	int height[] = {165, 170, 175, 180, 184};
	int height_len = 5;
	int k = 175;
	int ex = solution(height, height_len, k);	// 2
	printf("%d", ex);
}
