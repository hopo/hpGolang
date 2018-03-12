#include <stdio.h>

int findKim(char *names[]);

int main() {
	char *names1[] = {"Queen", "Tod", "Kim"};	// 2
	int ex1 = findKim(names1);
	printf("%d", ex1);

	char *names2[] = {"Gim", "Obama", "Xi"};	// -1
	int ex2 = findKim(names2);
	printf("\n%d", ex2);

	return 0;
}

int findKim(char *names[]) {
	int lnth = 3;
	for(int i = 0; i < lnth; ++i) {
		if(names[i] == "Kim") { return i; }
	}
	return -1;
}
