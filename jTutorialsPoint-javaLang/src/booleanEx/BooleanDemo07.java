package booleanEx;

import java.lang.*;

// String toString()
public class BooleanDemo07 {
	public static void main(String[] args) {
		
		// create 2 Boolean objects b1, b2
		Boolean b1, b2;

		// assign values to b1, b2
		b1 = new Boolean(true);
		b2 = new Boolean(null);
		
		// create 2 Strings s1, s2
		String s1, s2;
		
		// assing string value of objects b1, b2 to s1, s2
		s1 = b1.toString();
		s2 = b2.toString();
		
		String str1 = "String value of boolean object " +b1+ " is " +s1;
		String str2 = "String value of boolean object " +b2+ " is " +s2;
		
		// print s1, s2 vlues
		System.out.println(str1);
		System.out.println(str2);
	}
}
