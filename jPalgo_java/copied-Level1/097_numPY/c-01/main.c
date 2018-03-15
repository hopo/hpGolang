#include <stdio.h>

int numPY(char str[]);

int main() {
	int ex1 = numPY("pPoooyY"); // 1, true
	printf("%d", ex1);
	
	int ex2 = numPY("Pyy");     // 0, false
	printf("\n%d", ex2);
}

int numPY(char str[]) {
	int ppp, yyy, i;
	ppp = yyy = i = 0;

	while(str[i] != '\0') {
		switch(str[i]) {
			case 'p':
				ppp++;
				break;
			case 'y':
				yyy++;
				break;
			case 'P':
				ppp++;
				break;
			case 'Y':
				yyy++;
				break;
		}
		i++;
	}
	// printf("\nppp:%d, yyy:%d\n", ppp, yyy); // for check

	return (ppp == yyy)? 1 : 0;
}
