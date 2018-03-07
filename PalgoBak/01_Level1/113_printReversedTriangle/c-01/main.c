#include <stdio.h>

char *printReversedTriangle(int);

int main() {
	char* ex1 = printReversedTriangle(3); // ***\n**\n*
	printf("%s", ex1);

	return 0;
}

char *printReversedTriangle(int n) {
	char ret[100];
	int i, j, k, r = 0;

	for(i = 0; i < n; ++i) {
		k = n - i;	
		for(j = 0; j < k; ++j) {
			ret[r] = '*';
			r++;
		}
		if(k == 1) { break; }
		ret[r] = '\n';
		r++;
	}
	
	char box[r];
	for(i = 0; i < r; i++) { box[i] = ret[i]; } // choorigi
	box[r] = '\0'; // machimpyo
	char* dt = box;

	return dt;
}
