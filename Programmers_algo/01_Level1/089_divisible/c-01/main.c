#include <stdio.h>

int* divisible(int*, int);

int main() {
	int data[] = {5, 9, 7, 10};
	int n = 5;
	int* ex = divisible(data, n); // 5 10
	printf("%d %d", *ex, *(ex+1));

	return 0;
}

int* divisible(int* data, int n) {
	int i, j = 0, len = 0;
	
	for(i = 0; ; i++) { // int array find length??
		if(data[i] < 0) { break; }
		if(data[i] > 2 << 16) { break; }
		len++;
	}

	int ret[len];
	for(i = 0; i < len; i++) {
		if(!(*(data+i) % n)) {
			ret[j] = *(data+i);
			j++;
		}
	}
	
	int* dt = ret;

	return dt;
}
