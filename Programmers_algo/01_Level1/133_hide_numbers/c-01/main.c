#include <stdio.h>

char *hide_numbers(char pnum[]);

int main() {
	char *ex1;
	ex1 = hide_numbers("01033334444"); // "*******4444"
	printf("%s", ex1);

	char *ex2;
	ex2 = hide_numbers("027778888");   // "*****8888"
	printf("\n%s", ex2);

	return 0;
}

char *hide_numbers(char pnum[]) {
	int i, size = 0;
	for(i = 0; pnum[i] != '\0'; i++);
	size = i+1;

	char box[size];
	for(i = 0; i < size; i++) {
		if(i < size-5) { box[i] = '*'; }
		else { box[i] = pnum[i]; }
	}

	char ret[size];
	for(i = 0; i < size; i++) { ret[i] = box[i]; }	// choorigi
	ret[size] = '\0';	// machimpyo
	char *dt = ret;

	return dt;
}
