#include <stdio.h>

int *rm_small(int nums[], int lnth);
void display_iarr(int iarr[], int lnth);

int main() {
	int nums[] = {4, 3, 2, 1};
	int lnth = sizeof(nums)/sizeof(nums[0]);
	int *ex = rm_small(nums, lnth);	// [4 3 2]
	display_iarr(ex, 3);

	return 0;
}

int *rm_small(int nums[], int lnth) {
	int i, min;
	min = nums[lnth-1];

	for(i = 0; i < lnth; i++) {
		if(min > nums[i]) {
			min = nums[i];
		}
	}
	int count = 0;
	for(i = 0; i < lnth; i++) {
		if(min < nums[i]) {
			count++;
		}
	}

	int ret[count];
	int r = 0;
	for(i = 0; i < lnth; i++) {
		if(min < nums[i]) {
			ret[r] = nums[i];
			r++;
		}
	}
	
	int *dt;
	dt = ret;

	return dt;
}

void display_iarr(int iarr[], int lnth) {
	int box[lnth];
	int b = 0;
	for(int i = 0; i < lnth; i++) {
		box[b] = iarr[i];
		b++;
	}
	for(int i = 0; i < lnth; i++) {
		printf("%d ", box[i]);
	}
}
