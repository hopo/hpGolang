#include <stdio.h>

void toWeirdCase(char *);

int main() {
	char str1[] = "try hello world";
	toWeirdCase(str1);	// "TrY HeLlO WoRlD"
	printf("%s\n", str1);

	char str2[] = "coding is thinking not typing";
	toWeirdCase(str2); // "CoDiNg Is ThInKiNg NoT TyPiNg"
	printf("%s", str2);
	
	return 0;
}

void toWeirdCase(char *str) {
	int i = 0, w = 1;

	while(str[i] != '\0') {
		if(str[i] == ' ') {
			w = 1;	// reset w
			i++;	// next index
			continue;
		}

		if(w&1) { str[i] = str[i]-32; } // change upper char. 1st, 3rd, 5th...

		w++;
		i++; // next index
	}
}
