#include <stdio.h>

void display(int *arr);

int main() {
	int arr[] = {10, 8, 22};
	display(arr);
	
	/* find length. A or B */
	int size;
	size = *(&arr+1) - arr; 
	// size = sizeof(arr)/sizeof(arr[0]);
	printf("\nlenth: %d", size);

	return 0;
}

void display(int *arr) {
	printf("%d", arr[0]);
	printf("\n%d", arr[1]);
	printf("\n%d", arr[2]);
	printf("\n%d", arr[3]);
}

