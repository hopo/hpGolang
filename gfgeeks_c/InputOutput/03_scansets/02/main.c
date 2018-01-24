#include <stdio.h>

int main() {
	char str[128];

	printf("Enter a string: ");
	scanf("%[^o]s", str);

	printf("You entered: %s\n", str);

	return 0;
}

/*
Enter a string: http://geeks for geeks
You entered: http://geeks f
*/
