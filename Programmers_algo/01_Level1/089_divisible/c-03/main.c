#include <stdio.h>

/* int array of this program has a length data itself */

void hparr_conv(int lnth, int iarr[]);
void hparr_display(int hparr[]);
void divisible(int data[], int n);

int main() {
	int data[] = {5, 9, 7, 10}, n = 5;
	int lnth = sizeof(data)/sizeof(data[0]); 
	hparr_conv(lnth, data);
	divisible(data, n); // 5 10
	hparr_display(data);

	return 0;
}

void hparr_conv(int lnth, int iarr[]) {
	int i;
	int box[lnth+1];		

	// 0 index = length
	box[0] = lnth;

	for(i = 1; i < lnth+1; i++) {
		box[i] = iarr[i-1];	
	}

	// make hp array 
	for(i = 0; i < lnth+1; i++) {
		iarr[i] = box[i];
	}
}

void hparr_display(int hparr[]) {
	int i;
	int lnth = hparr[0]+1;
	for(i = 1; i < lnth; i++) {
		printf("%d ", hparr[i]);
	}
}

void divisible(int data[], int n) {
	int i, b = 0;
	int lnth = data[0]+1; // data[0] is length of origin array data

	int box[10];

	for(i = 1; i < lnth; i++) {
		if(!(data[i] % n)) {
			box[b] = data[i];
			b++;
		}
	}
	
	hparr_conv(b, box);
	
	for(i = 0; i < b+1; i++) {
		data[i] = box[i];
	}
}
