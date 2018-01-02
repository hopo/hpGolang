#include <stdio.h>

int* no_continous(char*);

int main() {
	char data1[] = "133303";	// 1 3 0 3
	int* ex1 = no_continous(data1);
	printf("%d %d %d %d", *ex1, *(ex1+1), *(ex1+2), *(ex1+3));

	printf("\n");

	char data2[] = "47330";  // 4 7 3 0
	int* ex2 = no_continous(data2);
	printf("%d %d %d %d", *ex2, *(ex2+1), *(ex2+2), *(ex2+3));

	return 0;
}

int* no_continous(char* data) {
	int i, j = 1, len = 0;
	for(i = 0; data[i] != '\0'; i++) { len++; }

	char box[len];
	box[0] = data[0];
	
	for(i = 1; i < len; i++) {
		if(data[i-1] != data[i]) {
			box[j] = data[i];
			j++;
		};
	}

	int ret[j];	
	for(i = 0; i < j; i++) { ret[i] = box[i]-'0'; }

	int* dt = ret;

	return dt;
}
