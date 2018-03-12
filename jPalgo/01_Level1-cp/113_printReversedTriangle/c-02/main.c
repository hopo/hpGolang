#include <stdio.h>
#include <stdlib.h>

void printReversedTriangle(int n, char ret[]);

int main() {
	int n = 5;
	char ret[100];
	printReversedTriangle(n, ret); // ***\n**\n*
	printf("%s", ret);

	return 0;
}

void printReversedTriangle(int n, char ret[]) {
	if(n > 13) {
		printf("num is over.");
		exit(0);
	}
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
	ret[r] = '\0';
}
