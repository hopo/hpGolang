package level01;

import java.lang.String;

public class StringMiddle {
	public static void main(String[] args) {
		String str1 = new String("power");
		String ex1 = stringmiddle(str1); // w
		System.out.println(ex1);

		String str2 = new String("java");
		String ex2 = stringmiddle(str2); // av
		System.out.print(ex2);
	}

	public static String stringmiddle(String str) {
        char[] carr = str.toCharArray();
		int lnth = carr.length;
		int lnth_half = lnth/2;

        String ret = "";
        ret += carr[lnth_half];
		if(lnth%2 == 0) {
            ret = carr[lnth_half-1] + ret;
		}

		return ret;
	}
}
