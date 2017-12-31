#include <stdio.h>
#include "/Users/HPMBA/workspace/hinclude/chstd.h"

int iarrlen(int *iarr) {
	int length = 0;
	while(1) {
		printf("%d, ", iarr[length]);
		if(iarr[length] == 0) break; // have to modify
		else if(iarr[length] < 0) break; // have to modify
		else if(iarr[length] > (2<<16)) break; // have to modify
		length++;
	}
	return length;
}

int main() {
	int iarr[] = {2};

	printf("\n%d", iarrlen(&iarr));
	
	return 0;
}
