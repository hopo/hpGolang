package langbyte;

public class ValueOfExam {

	public static void main(String[] args) {
		// create a String s and assign value to it
		String s = "+120";
		
		// create a Byte object b
		Byte b;
		
		/**
		 *  static method is called using class name.
		 *   assign Byte instance value of s to b
		 */
		b = Byte.valueOf(s);
		
		String str = "Byte value of string " + s + " is " + b;
		
		// print b value
		System.out.println(str);
	}
}
