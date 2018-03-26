package langcharacter;

public class CodePointAt {
	public static void main(String[] args) {
		// create a char array c and assign values
		char[] c = new char[] {'a', 'b', 'c', 'd', 'e'};
		
		// create 2 int's res, index and assign value to index
		int res, index = 0;
		
		// assign result of codPointAt on array c at index to res
		res = Character.codePointAt(c, index);
		
		String str = "Unicode code point is " + res;
		
		// print res value
		System.out.println(str);
	}
}
