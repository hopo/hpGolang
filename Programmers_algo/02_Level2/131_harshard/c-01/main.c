#include <stdio.h>

int harshad(int);

int main() {
	int num1 = 18, ex1;
	ex1 = harshad(num1);	// 1, 1+8=9, 18%9 == 0
	printf("%d\n", ex1);

	int num2 = 13, ex2;
	ex2 = harshad(num2);	// 0, 1+3=4, 13%4 == 1
	printf("%d\n", ex2);

	int num3 = 342, ex3;
	ex3 = harshad(num3);	// 1, , 3+4+2=9 342%9 == 0
	printf("%d", ex3);

	return 0;
}

int harshad(int num) {
	int n, dv = 10, smpl = 0, ret = 0;
	n = num;

	/* jaritsu of num */
	while(1) {
		if(num-dv < 0) {
			dv /= 10;
			break;
		}
		dv *= 10;
	}

	/* total each num */
	while(1) {
		if(!dv) {
			break;
		}
		smpl += num/dv;
		num %= dv;
		dv /= 10;
	}
	
	/* does have 'n%smpl' remainder */
	ret = (n%smpl)? 0 : 1;

	return ret;
}
