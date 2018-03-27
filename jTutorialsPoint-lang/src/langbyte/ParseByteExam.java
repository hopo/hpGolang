package langbyte;

public class ParseByteExam {
	public static void main(String[] args) {
		// create 2 byte primitives bt1, bt2
		byte bt1, bt2;
		
		// create and assign values to String's s1, s2
		String s1 = "+123";
		String s2 = "-123";
		
		/**
		 *  static method is called using class name.
		 *   assign parseByte result on s1, s2 to bt1, bt2
		 */
		bt1 = Byte.parseByte(s1);
		bt2 = Byte.parseByte(s2);
		
		String str1 = "Parse byte value of " + s1 + " is " + bt1;
		String str2 = "Parse byte value of " + s2 + " is " + bt2;
		
		// print bt1, bt2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}