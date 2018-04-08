package level02;

import java.util.Arrays;

public class DigitReverse {
	public static void main (String[] args) {
		int dig1 = 12345;
		int[] ex1 = dreverse(dig1); // [5, 4, 3, 2, 1]
		System.out.println(Arrays.toString(ex1));

		int dig2 = 321;
		int[] ex2 = dreverse(dig2); // [1, 2, 3]
		System.out.println(Arrays.toString(ex2));
				
	}
	
	static int[] dreverse(int digit) {
		// int digit convert char Array type.
		char[] cStr = Integer.toString(digit).toCharArray();

		int lnth = cStr.length;
		int[] ret = new int[lnth];
		
		// reverse sort and indexing int[] ret.
		for(int i = 0; i < lnth; i++) {
			ret[i] = cStr[lnth-1-i] - 48;
		}
		
		return ret;
	}
}
