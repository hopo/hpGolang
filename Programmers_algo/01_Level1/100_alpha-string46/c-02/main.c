#include <stdio.h>

int alpha_string46(char str[]);

int main() {
	int ex1 = alpha_string46("a234"); // 0, False, "string 'a'"
	printf("%d", ex1);

	int ex2 = alpha_string46("1234"); // 1, True, "all int"
	printf("\n%d", ex2);
	
	return 0;
}

int alpha_string46(char str[]) {
	int i = 0;
	while(str[i] != '\0') {
		if(str[i] < '0' || str[i] > '9') { return 0; }
		i++;
	}
	
	return 1;
}
