#include <stdio.h>
#include <string.h>

typedef struct Namecard
{
	char name[50];
	int number;
} Ncard;

void sort_dictionary(Ncard cards[], int lnth);
void display_cardsArray(Ncard cards[], int lnth);

int main() {
	Ncard n1, n2, n3;
	strcpy(n1.name, "김철수"); n1.number = 78;  
	strcpy(n2.name, "이하나"); n2.number = 97;
	strcpy(n3.name, "정진원"); n3.number = 88;
	Ncard cards[] = {n1, n2, n3};
	int lnth = sizeof(cards)/sizeof(cards[0]);
	sort_dictionary(cards, lnth);
	display_cardsArray(cards, lnth);

	return 0;
}

void sort_dictionary(Ncard cards[], int lnth) {
	int i, j;
	int cmp; 
	Ncard temp;

	for(i = 0; i < lnth-1 ; i++) {
		for(j = 1; j < lnth ; j++) {
			cmp = strcmp(cards[i].name, cards[j].name);
			if(0 < cmp) {
				temp = cards[i];
				cards[i] = cards[j];
				cards[j] = temp;
			}
		}
	}
}

void display_cardsArray(Ncard cards[], int lnth) {
	for(int i = 0; i < lnth; i++) {
		printf("(%s, %d) ", cards[i].name, cards[i].number);
	}
}
