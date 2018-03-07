// https://programmers.co.kr/learn/courses/34/lessons/1867

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int* func_a(int arr[], int arr_size, int num){
    int* ret = (int*)malloc(sizeof(int)*(arr_size - 1));
    int idx = 0;
    for(int i = 0; i < arr_size; ++i)
        if(arr[i] != num)
            ret[idx++] = arr[i];
    return ret;
}

int func_b(int a, int b){
    if(a >= b)
        return a - b;
    else
        return b - a;
}

int func_c(int arr[], int arr_size){
    int ret = -1;
    for(int i = 0; i < arr_size; ++i)
        if(ret<arr[i])
            ret = arr[i];
    return ret;
}

int solution(int visitor[], int n ) {
    int max_first = func_ /**/ ;
    int* visitor_removed = func_ /**/ ;
    int max_second = func_ /**/ ;
    int answer = func_ /**/ ;
    return answer;
}
