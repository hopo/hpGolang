#include <stdio.h>
#include <string.h>

typedef char *string;
 
int strToInt(string);
int convctoi(char);

int main() {
	int ex1 = strToInt("-1234");	// -1234
	printf("%d", ex1);

	int ex2 = strToInt("392");	// 392
	printf("\n%d", ex2);

	return 0;
}

int strToInt(string s) {
	int lnth = strlen(s);
	int ret = 0, i, j, p = 1;
	for(i = 0; i < lnth; i++) {
		j = lnth-1-i;
		if('0' <= s[j] && s[j] <= '9') { 
			ret += convctoi(s[j])*p;
		}
		p *= 10;
	}

	return (s[0] == '-')? ret *= -1 : ret;
}

int convctoi(char c) {
	return (int)(c-48);	
}
