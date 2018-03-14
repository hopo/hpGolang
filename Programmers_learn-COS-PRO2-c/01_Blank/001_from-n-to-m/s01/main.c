#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int func_a(int k){
    int sum = 0;
    for(int i = 0; i < k ; ++i)
        sum += i+1;
    return sum;
}

int solution(int n, int m) {
    int sum_to_m = func_a(m);
    int sum_to_n = func_a(n-1);
    int answer = sum_to_m - sum_to_n;
    return answer;
}

int main() {
	int ex1 = solution(5, 10);	// 45
	printf("%d\n", ex1);

	int ex2 = solution(6, 6);	// 6
	printf("%d", ex2);
}
