package langboolean;

// static boolean parseBoolean(String s)
public class ParseBooleanExam {
	public static void main(String[] args) {
		
		// create and assign values to String;s s1, s2
		String s1 = "TRue";
		String s2 = "yes";
		
		// create 2 boolean primitives bool1, bool2
		boolean bool1, bool2;
		
		/**
		 *  static method is called using class name
		 *  apply result of parseBoolean on s1, s2 to bool1, bool2
		 */
		bool1 = Boolean.parseBoolean(s1);
		bool2 = Boolean.parseBoolean(s2);
		
		String str1 = "Parse boolean on " +s1+ " gives " +bool1;
		String str2 = "Parse boolean on " +s2+ " gives " +bool2;
		
		// print b1, b2 balues
		System.out.println(str1);
		System.out.println(str2);
	}
}
