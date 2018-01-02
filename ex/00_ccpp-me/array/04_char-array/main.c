// ..ing


#include <stdio.h>

void modi(char* words[]) {
	int i = 0;
	printf("before:\n");
	while(i < 5) {
		printf("%s ", *(words+i));
		i++;
	}

	char* ret[] = {*(words+4), *(words+3), *(words+1)};

	int j = 0;
	printf("\nafter:\n");
	while(j < 3) {
		printf("%s ", *(ret+j));
		j++;
	}
}

int main() {
	char* words[] = {"sun", "bed", "car", "exo", "you"};
	modi(words);

	return 0;
}

