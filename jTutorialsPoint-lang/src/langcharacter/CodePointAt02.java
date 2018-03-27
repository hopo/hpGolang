package langcharacter;

// ChaSequence obj
public class CodePointAt02 {
	public static void main(String[] args) {
		// create a CharSequence seq and assign value
		CharSequence seq = "Hello";
		
		// create and assign value to index
		int index = 4;
		
		// create an int res
		int res;
		
		// assign result of codePointAt in seq at index to res
		res = Character.codePointAt(seq, index); // 'o'
		
		String str = "Unicode code point is " + res;
		
		// print res value
		System.out.println(str);
	}
}
