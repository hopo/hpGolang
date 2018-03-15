#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int* solution(int scores[], int scores_len) {
    int* grade_counter = (int*)malloc(sizeof(int)*5);
    for(int i = 0; i < 5; ++i)
        grade_counter[i] = 0;

    for(int i = 0; i < scores_len; ++i)
    {
        if(scores[i] > 84)
            grade_counter[0] += 1;
        else if(scores[i] > 69)
            grade_counter[1] += 1;
        else if(scores[i] > 54)
            grade_counter[2] += 1;
        else if(scores[i] > 39)
            grade_counter[3] += 1;
        else
            grade_counter[4] += 1;
    }
    return grade_counter;
}

int main() {
	int scores[] = {86, 72, 98, 60, 45};
	int scores_len = 5;
	int* ex = solution(scores, scores_len);	// {2, 1, 1, 1, 0}
	for(int i = 0; i < scores_len; i++) { printf("%d ", ex[i]); }
}
