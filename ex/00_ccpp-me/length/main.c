#include <stdio.h>
#include "/Users/HPMBA/workspace/hinclude/chstd.h"

int iarrlen(int *iarr) {
	printf("!!! %p\n", iarr);
	printf("!!! %p\n", iarr+1);
	printf("!!! %p\n", iarr+2);
	printf("!!! %p\n", iarr+3);

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
	int iarr[] = {2, 4, 6, 8};

	printf("\n%d", iarrlen(&iarr));
	
	return 0;
}
