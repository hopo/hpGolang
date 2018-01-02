#include <stdio.h>

void prnt(char* words[]);

int main() {
	char* words[] = {"sun", "bed", "car", "exo", "you"};

	prnt(words);

	return 0;
}

void prnt(char* words[]) {
	int i = 0;
	while(i < 5) {
		printf("%s ", *words);
		words++;
		i++;
	}
}
