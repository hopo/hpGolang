#include <stdio.h>

int getMinMax(int *, int *);

int main() {
	int arra1[2] = {1, 2}, arrb1[2] = {3, 4}, ex1;
	ex1 = getMinMax(arra1, arrb1);	// 10
	printf("%d\n", ex1);
	
	int arra2[2] = {4, 2}, arrb2[2] = {3, 9}, ex2;
	ex2 = getMinMax(arra2, arrb2);	// 30
	printf("%d\n", ex2);

	int arra3[2] = {2, 3}, arrb3[2] = {5, 7}, ex3;
	ex3 = getMinMax(arra3, arrb3);	// 29
	printf("%d", ex3);

	return 0;
}

int getMinMax(int *arra, int *arrb) {
	int i, j, lnth = 2, min, l = 0, data[4];
	/* multiplicate each elements */
	for(i = 0; i < lnth; i++) {
		for(j = 0; j < lnth; j++) {
			data[l] = arra[i] * arrb[j];
			l++;
		}
	}	

	/* find min value. comparing sum of each multiplcated value */
	min = data[0] + data[l-1];
	for(i = 0; i < l/2; i++) {
		if(min > data[i] + data[l-1-i]) {
			min = data[i] + data[l-1-i];
		}
	}
	
	return min;
}
