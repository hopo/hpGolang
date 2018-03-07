#include <stdio.h>

void caesar(char *, int);

int main() {
	char key1[] = "a B z";
	int n1 = 4;
	caesar(key1, n1);   // "e F d"
	printf("%s\n", key1);

	char key2[] = "A b Z g";	
	int n2 = 5;
	caesar(key2, n2); // "F g E l"
	printf("%s", key2);
}

void caesar(char *key, int n2) {
	int i = 0;

	while(key[i] != '\0') {
		if(key[i] != ' ') {
			if(key[i] == 'Z') { key[i] = 'A'-1; }
			if(key[i] == 'z') { key[i] = 'a'-1; }
			key[i] += n2;
		}
		i++;
	}
}
