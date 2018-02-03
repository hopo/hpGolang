#include <stdio.h>

int collats(int);

int main() {
	int num1 = 6, ex1;
	ex1 = collats(num1);	// 8, 6:3-10-5-16-8-4-2-1
	printf("%d\n", ex1);

	int num2 = 11, ex2;
	ex2 = collats(num2);	// 14
	printf("%d", ex2); 
	
	return 0;
}

int collats(int num) {
	int count = 0;

	while(num != 1) {
		if(num&1) { num = num*3+1; }
		else { num /=2; }
		count++;
	}

	return count;
}

