#include <stdio.h>

int nlcm(int*, int);

int main() {
	int nums1[] = {3, 5, 6, 7, 8, 9, 15}, nums_len1, ex1;
	nums_len1 = sizeof(nums1)/sizeof(nums1[0]);
	ex1 = nlcm(nums1, nums_len1);	// 2520
	printf("%d\n", ex1);

	int nums2[] = {2, 6, 8, 14}, nums_len2, ex2;
	nums_len2 = sizeof(nums2)/sizeof(nums2[0]);
	ex2 = nlcm(nums2, nums_len2);	// 168
	printf("%d", ex2);

	return 0;
}

int nlcm(int* nums, int nums_len) {
	int i, mi = 1, count = 0, lcm ;
	int max = nums[0];
			
	for(i = 1; i < nums_len; i++) {
		if(nums[i] > max) { max = nums[i]; }
	}
	
	while(count != nums_len) {
		lcm = max * mi;
		for(i = 0; i < nums_len; i++) {
			if(lcm%nums[i]) {
				count = 0;
				break;
			}
			count++;
		}
		mi++;
	}

	return lcm;
}
