#include <stdio.h>


int main() {
	int a, b, c;

	printf("Enter value of a in decimal format: ");
	scanf("%d", &a);

	printf("Enter value of a in octal format: ");
	scanf("%i", &b);

	printf("Enter value of a in hexadecimal format: ");
	scanf("%i", &c);

	printf("a = %i, b = %i, c = %i", a, b, c);

	return 0;
}

/*
Enter value of a in decimal format: 12
Enter value of a in octal format: 012
Enter value of a in hexadecimal format: 0x12
a = 12, b = 10, c = 18
*/
