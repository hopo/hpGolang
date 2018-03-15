#include <stdio.h>

void strange_sort(char *words[], int n, int lnth);
void display_cstrArray(char *words[], int lnth);

int main() {
	char *words[] = {"sun", "bed", "car"}; // car, bed, sun
	int n = 1;
	int lnth = sizeof(words)/sizeof(words[0]);

	strange_sort(words, n, lnth);
	display_cstrArray(words, lnth);

	return 0;
}

void strange_sort(char *words[], int n, int lnth) {
	int i, j;
	char *temp;
	
	for(i = 0; i < lnth-1; i++) {
		for(j = 1; j < lnth; j++) {
			if(words[i][n] > words[j][n]) {
				temp = words[i];
				words[i] = words[j];
				words[j] = temp;
			}
		}
	}
}

void display_cstrArray(char *words[], int lnth) {
	for(int i = 0; i < lnth; i++) {
		printf("%s ", words[i]);
	}
}
