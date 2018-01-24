#include <stdio.h>

int main() {
	char str[128];

	printf("Enter a string: ");
	scanf("%[A-Z]s", str);

	printf("You entered: %s\n", str);

	return 0;
}

/*
Enter a string: GEEKs_for_geeks
You entered: GEEK
 */
