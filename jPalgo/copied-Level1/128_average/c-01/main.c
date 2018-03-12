#include <stdio.h>

float average(float arrf[], int size);

int main() {
	float arrf[] = {5, 3, 4};	//4, 12/3
	float ex1 = average(arrf, sizeof(arrf)/sizeof(arrf[0]));
	printf("%f", ex1);
}

float average(float arrf[], int size) {
	int i, sum = 0;
	for(i = 0; i < size; i++) {
		sum += arrf[i];
	}
	return (sum/size);
}
