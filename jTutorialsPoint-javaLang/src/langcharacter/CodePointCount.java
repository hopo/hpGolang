package langcharacter;

public class CodePointCount {
	public static void main(String[] args) {
		// create a char array c and assign values
		char[] c = new char[] {'a', 'b', 'c', 'd', 'e'}	;
		
		// create and assign value to offset, count
		int offset = 1, count = 3;
		
		// create an int res
		int res;
		
		// assign result of codePointCount on subarray of c to res
		res = Character.codePointCount(c, offset, count);
		
		String str = "No. of Unicode code points is " + res;
		
		// print res value
		System.out.println(str);
	}

}
