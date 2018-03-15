#include <stdio.h>
#include <string.h>

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

	char box[lnth+1];
	for(i = 0; i < lnth; i++) {
		j = lnth-1-i;
		box[i] = str[j];
	}

	strcpy(str, box);
}
