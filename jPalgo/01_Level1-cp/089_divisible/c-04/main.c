#include <stdio.h>

/* int array(*hIarr) of this program has a length data itself at index 0 */
typedef int *hIarr;

hIarr hIarr_conv(int lnth, int arr[]);
hIarr divisible(hIarr dt, int n);
void hIarr_display(hIarr dt);

int main() {
	int data[] = {5, 9, 7, 10}, n = 5;
	int lnth = sizeof(data)/sizeof(data[0]); 
	hIarr dt = hIarr_conv(lnth, data);
	hIarr ex = divisible(dt, n); // 5 10
	hIarr_display(ex);

	return 0;
}

hIarr hIarr_conv(int lnth, int arr[]) {
	int i;
	hIarr box;	

	box[0] = lnth;
	for(i = 1; i < lnth+1; i++) {
		box[i] = arr[i-1];
	}
	
	return box;	
}

hIarr divisible(hIarr dt, int n) {
	int i, r = 1;
	int lnth = dt[0]+1;

	hIarr ret;

	for(i = 1; i < lnth; i++) {
		if(!(dt[i] % n)) {
			ret[r] = dt[i];
			r++;
		}
	}
	ret[0] = r-1;

	return ret;
}

void hIarr_display(hIarr dt) {
	int i;
	int lnth = dt[0]+1;
	for(i = 1; i < lnth; i++) {
		printf("%d ", dt[i]);
	}
}
