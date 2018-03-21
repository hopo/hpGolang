package level01;

import java.util.Arrays;

public class ReverseStr {
	public static void main(String[] args) {
		ReverseStr rs = new ReverseStr();
		System.out.println( rs.reverseStr("Zbcdefg") ); // "gfedcbZ"
	}

	public String reverseStr(String str){
        String ret = "";
        char[] ctmp = str.toCharArray();
        Arrays.sort(ctmp);

        for(int i = 0; i < ctmp.length; i++) {
            ret = Character.toString(ctmp[i]) + ret;
        }

        return ret;
	}

}
