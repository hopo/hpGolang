#include <stdio.h>

void reverseStr(char str[]);

int main() {
	char str[] = "Zbcdefg";
	reverseStr(str); // "gfedcbZ"
	printf("%s", str);

	return 0;
}

void reverseStr(char str[]) {
	int i, j, lnth;
	for(i = 0; str[i] != '\0'; i++);
	lnth = i;

	char temp;
	for(i = 0; i < lnth/2; i++) {
		j = lnth-1-i;
		temp = str[i];
		str[i] = str[j];
		str[j] = temp;
	}
}
