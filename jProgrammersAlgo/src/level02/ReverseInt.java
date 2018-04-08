package level02;

import java.util.Arrays;

public class ReverseInt {
	public static void main(String[] args) {
		int iNum1 = 118372;
		int ex1 = reverseint(iNum1); // 873211
		System.out.println(ex1);
	}
	
	static int reverseint(int num) {
		// int convert to char array. and sorting
		char[] chArr = Integer.toString(num).toCharArray();
		Arrays.sort(chArr);

		// make return int
		int ret = chArr[0]-48;
		for(int i = 1; i < chArr.length; i++) {
			ret += (chArr[i]-48) * (Math.pow(10, i));
		}
		
		return ret;
	}

}
