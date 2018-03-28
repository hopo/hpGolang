package level01;

import java.util.Arrays;

public class GcdLcm {
	public static void main(String[] args) {
		int a1 = 3, b1 = 12;
		int[] ex1 = gcdlcm(a1, b1);	// [3, 12]
		System.out.println(Arrays.toString(ex1));

		int a2 = 4, b2 = 7;
		int[] ex2 = gcdlcm(a2, b2);	// [1, 28]
		System.out.println(Arrays.toString(ex2));
	}

	static int[] gcdlcm(int a, int b) {
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