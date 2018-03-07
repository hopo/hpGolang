// https://programmers.co.kr/learn/courses/34/lessons/1865

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int func_a(int k){
    int sum = 0;
    for(int i = 0; i /**/ ; ++i)
        sum += /**/;
    return sum;
}

int solution(int n, int m) {
    int sum_to_m = func_a(m);
    int sum_to_n = func_a(n-1);
    int answer = sum_to_m - sum_to_n;
    return answer;
}
