#include <stdio.h>
#include <string.h>

typedef struct Namecard
{
	char name[10];
	int number;
} Ncard;

int main() {
	Ncard n1, n2, n3;
	strcpy(n1.name, "김철수"); n1.number = 78;  
	strcpy(n2.name, "이하나"); n2.number = 97;
	strcpy(n3.name, "정진원"); n3.number = 88;
	
	Ncard cards[] = {n1, n2, n3};

	int x = sizeof(cards);
	int y = sizeof(n1);
	int lnth = x/y;

	printf("sizeof each: %d, %d\n", x, y);
	printf("length: %d\n", lnth);

	for(int i = 0; i < 3; i++) {
		printf("%s, ", cards[i].name);
		printf("%d\n", cards[i].number);
	}

	return 0;
}
