#include <stdio.h>

int noOvertime(int, int, int *);

int main() {
	int n1 = 4, lnth1, wks1[] = {4, 3, 3}, ex1;
	lnth1 = sizeof(wks1)/sizeof(wks1[0]);
	ex1 = noOvertime(n1, lnth1, wks1);	// 12, [2, 2, 2] -> (2**2) *3
	printf("%d\n", ex1);

	int n2 = 7, lnth2, wks2[] = {9, 11}, ex2;
	lnth2 = sizeof(wks2)/sizeof(wks2[0]);
	ex2 = noOvertime(n2, lnth2, wks2);	// 85, [7, 6] -> ((6+1)**2)) + (6**2)
	printf("%d", ex2);
}

int noOvertime(int n, int lnth, int *wks) {
	int i, twks = 0, at, rt, ret = 0;

	/* total work time */
	for(i = 0; i < lnth; i++) { twks += wks[i]; }

	/* average time and remainder time */
	at = (twks-n)/lnth;
	rt = (twks-n)%lnth;

	/* (at**2) + (at**2) + ... (separating rt value) */
	for(i = 0; i < lnth; i++) {
		if(!rt) { ret += at*at; }
		else if(rt) {
			ret += (at+1) * (at+1);
			rt--;
		}
	}

	return ret;
}


/*
func noOvertime(n int, works []int) int {
	// * (total works - n) / wparts ... remainder
	var total int
	wparts := len(works)
	for _, v := range works {
		total += v
	}
	ovtime := total - n
	div := ovtime / wparts
	rmnd := ovtime % wparts

	var ret, q int
	for i := 0; i < wparts; i++ {
		q = 0
		if rmnd != 0 {
			q = 1
			rmnd--
		}
		ret += (div + q) * (div + q)
	}

	return ret
}

func least_ot(n int, works []int) {
	var total int
	//wparts := len(works)
	for _, v := range works {
		total += v
	}
	//ov := total - n

	var ret int
	var x, y, z []int
	for i := 0; i < 7; i++ {
		x = []int{6, 5, 4, 3, 4, 3, 2}
		y = []int{0, 1, 2, 3, 1, 2, 2}
		z = []int{0, 0, 0, 0, 1, 1, 2}
		ret = (x[i] * x[i]) + (y[i] * y[i]) + (z[i] * z[i])
		fmt.Println(ret)
	}
}
*/
