package level01;

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
