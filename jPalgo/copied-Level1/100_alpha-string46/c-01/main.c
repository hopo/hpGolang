#include <stdio.h>
#include <ctype.h>

int alpha_string46(char str[]);

int main() {
	char str1[] = "a234";
	int ex1 = alpha_string46(str1); // 0, False, "string 'a'"
	printf("%d", ex1);

	char str2[] = "1234";
	int ex2 = alpha_string46(str2); // 1, True, "all int"
	printf("\n%d", ex2);
	
	return 0;
}

int alpha_string46(char str[]) {
	int i, d;
	for(i = 0; str[i] != '\0'; i++) {
		d = isdigit(str[i]);
		if(d == 0) { return 0; }
	}

	return 1;
}
