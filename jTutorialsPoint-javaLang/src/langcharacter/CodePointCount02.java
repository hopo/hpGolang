package langcharacter;

public class CodePointCount02 {
	public static void main(String[] args) {
		// create a CharSequence seq and assign value
		CharSequence seq = "Hello World!";
		
		// create and assign value to bi, ei
		int bi = 4, ei = 8;
		
		// create an int res
		int res;
		
		// assign result of codePointCount on seq to res using bi, ei
		res = Character.codePointCount(seq, bi, ei);
		
		String str = "No. of Unicode code points is " + res;
		
		// print res value
		System.out.println(str);
	}
}
