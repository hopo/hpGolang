// https://programmers.co.kr/learn/courses/34/lessons/1866

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int func_a(int s[], int arr_size){
    int ret = 0;
    for(int i = 0; i < arr_size; ++i)
        if(s[i] > ret)
            ret = s[i];
    return ret;
}

int func_b(int s[], int arr_size){
    int ret = 0;
    for(int i = 0; i < arr_size; ++i)
        ret += s[i];
    return ret;
}

int func_c(int s[], int arr_size){
    int ret = 101;
    for(int i = 0; i < arr_size; ++i)
        if(s[i] < ret)
            ret = s[i];
    return ret;
}

int solution(int scores[], int scores_len) {
    int sum = func_ /**/ ;
    int max_score = func_ /**/ ;
    int min_score = func_ /**/ ;
    return sum - max_score - min_score;
}
