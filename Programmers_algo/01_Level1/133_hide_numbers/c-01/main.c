#include <stdio.h>

char *hide_numbers(char pnum[]);

int main() {
	printf("\b"); // why???

	char *ex1;
	ex1 = hide_numbers("01033334444"); // "*******4444"
	printf("%s", ex1);

	char *ex2;
	ex2 = hide_numbers("027778888");   // "*****8888"
	printf("\n%s", ex2);

	return 0;
}

char *hide_numbers(char pnum[]) {
	int i, lnth = 0;
	for(i = 0; pnum[i] != '\0'; i++);
	lnth = i;

	char ret[lnth];
	for(i = 0; i < lnth; i++) {
		if(i < lnth-4) { ret[i] = '*'; }
		else { ret[i] = pnum[i]; }
	}
	ret[lnth] = '\0';

	char *dt;
	dt = ret;

	return dt;
}
