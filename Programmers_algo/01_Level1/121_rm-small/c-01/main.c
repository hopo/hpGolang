#include <stdio.h>

int *rm_small(int *nums);

int main() {
	int nums1[] = {4, 3, 2, 1};
	int* ex1 = rm_small(nums1); // [4 3 2]
	printf("%d %d %d", ex1[0], ex1[1], ex1[2]);

	return 0;
}

int *rm_small(int *nums) {
	int i, lnth, min;
	min = nums[0];

	for(i = 1; nums[i] < 2<<24; i++) { // we need find int length! A)c have no idea.
		if(nums[i] < 0) { break; }
		if(min > nums[i]) { min = nums[i]; }	
	}
	lnth = i;

	int ret[lnth-1], r = 0;
	for(i = 0; i < lnth; i++) {
		if(nums[i] != min) {
			ret[r] = nums[i];	
			r++;
		}
	}
	
	int* dt;
	dt = ret;

	return dt;
}

