#include <stdio.h>

int main() {
	char str[128];

	printf("Enter a string with spaces: ");
	scanf("%[^\n]s", str);

	printf("You entered: %s\n", str);

	return 0;
}

/*
Enter a string with spaces: Geeks For Geeks
You entered: Geeks For Geeks 
*/
