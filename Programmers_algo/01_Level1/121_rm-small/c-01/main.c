// ing

#include <stdio.h>

void rm_small(int*);

int main() {
	int nums1[] = {4, 3, 2, 1};
	rm_small(nums1); // [4 3 2]

	//int nums2[] = {10, 8, 22};
	//rm_small(nums2)  // [10 22]

	return 0;
}

void rm_small(int* nums) {
	int i, lnth, min;
	min = nums[0];

	for(i = 1; nums[i] < 2<<24; i++) {
		if(min > nums[i]) { min = nums[i]; }	
	}
	lnth = i;

}
