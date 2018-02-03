#include <stdio.h>

void hide_numbers(char *pnum);

int main() {
	char pnum1[] = "01033334444";
	hide_numbers(pnum1); // "*******4444"
	printf("%s", pnum1);

	char pnum2[] = "027778888";
	hide_numbers(pnum2);   // "*****8888"
	printf("\n%s", pnum2);

	return 0;
}

void hide_numbers(char *pnum) {
	int i, lnth = 0;
	for(i = 0; pnum[i] != '\0'; i++);
	lnth = i;

	for(i = 0; i < lnth-4; i++) { pnum[i] = '*'; }
}
