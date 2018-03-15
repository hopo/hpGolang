// package level01;

import java.util.Arrays;

public class GcdLcm {
	public static void main(String[] args) {
		GcdLcm gl1 = new GcdLcm();
		int[] ex1 = gl1.gcdlcm(3, 12);	// 3, 12
		System.out.println(Arrays.toString(ex1));

		GcdLcm gl2 = new GcdLcm();
		int[] ex2 = gl2.gcdlcm(4, 7);	// 1, 28
		System.out.println(Arrays.toString(ex2));
	}

	int[] gcdlcm(int a, int b) {
		int tmp;
		if(a > b) {
			tmp = a;
			a = b;
			b = tmp;
		}

		int mn = a, mx = b;
		while(mn > 0) {
			tmp = mx%mn;
			mx = mn;
			mn = tmp;
		}

		int[] ret = {mx, (a*b)/mx};
		return ret;
	}
}


/*
int gcd(int n1, int n2);
int lcm(int n1, int n2);
int* gcdlcm(int n1, int n2);


int gcd(int n1, int n2) {
	// bigger - smaller ?
	while(n1 != n2) {
		if(n1 > n2) { n1 -= n2; }
		else { n2 -= n1; }
	} 

	return n1;
}

int lcm(int n1, int n2) {
	int max;

	// maximum value between n1 and n2 is stored in max
	max = (n1 > n2)? n1 : n2;

	do {
		if(max%n1 == 0 && max%n2 == 0) { break; }
		else { ++max; }
	}
	while(1);

	return max;
}

int* gcdlcm(int n1, int n2) {
	int r1, r2;
	r1 = gcd(n1, n2);
	r2 = lcm(n1, n2);

	int ret[] = {r1, r2};

	int* dt = ret;

	return dt;
}

*/
