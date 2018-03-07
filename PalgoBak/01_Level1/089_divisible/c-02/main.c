#include <stdio.h>

void divisible(int data[], int n, int lnth);
void display(int iarr[], int lnth);

int main() {
	int data[] = {5, 9, 7, 10};
	int lnth = sizeof(data)/sizeof(data[0]);
	int n = 5;
	divisible(data, n, lnth); // 5 10
	display(data, 2); // 2 is disible.b

	return 0;
}

void divisible(int data[], int n, int lnth) {
	int i, b = 0;
	int box[lnth];

	for(i = 0; i < lnth; i++) {
		if(!(data[i] % n)) {
			box[b] = data[i];
			b++;
		}
	}

	// input data
	for(i = 0; i < b; i++) {
		data[i] = box[i];
	}
}

void display(int iarr[], int lnth) {
	int i;
	for(i = 0; i < lnth; i++) {
		printf("%d ", iarr[i]);
	}
}
