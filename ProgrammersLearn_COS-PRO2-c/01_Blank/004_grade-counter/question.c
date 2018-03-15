// https://programmers.co.kr/learn/courses/34/lessons/1868

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int* solution(int scores[], int scores_len) {
    int* grade_counter = (int*)malloc(sizeof(int)*5);
    for(int i = 0; i < 5; ++i)
        grade_counter[i] = 0;

    for(int i = 0; i < scores_len; ++i)
    {
        if( /**/ )
            grade_counter[0] += 1;
        else if( /**/ )
            grade_counter[1] += 1;
        else if( /**/ )
            grade_counter[2] += 1;
        else if( /**/ )
            grade_counter[3] += 1;
        else
            grade_counter[4] += 1;
    }
    return grade_counter;
}
