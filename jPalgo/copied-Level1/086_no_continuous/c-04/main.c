#include <stdio.h>
#include <string.h>

void no_continous(char data[]);
void display(char data[]);

int main() {
	char data1[] = "133303";	// 1 3 0 3
	no_continous(data1);
	display(data1);

	printf("\n");
	char data2[] = "47330";  // 4 7 3 0
	no_continous(data2);
	display(data2);

	return 0;
}

void no_continous(char data[]) {
	int lnth, i, j = 1;
	lnth = strlen(data);

	char box[lnth+1];
	box[0] = data[0];
	
	for(i = 1; i < lnth; i++) {
		if(data[i-1] != data[i]) {
			box[j] = data[i];
			j++;
		}
	}
	box[j] = '\0';
	
	strcpy(data, box);
}

void display(char data[]) {
	int i;
	for(i = 0; data[i] != '\0'; i++) {
		printf("%c ", data[i]);
	}
}
