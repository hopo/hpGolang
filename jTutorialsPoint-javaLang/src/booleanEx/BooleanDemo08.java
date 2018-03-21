package booleanEx;

import java.lang.*;

// static String toString(boolean b)
public class BooleanDemo08 {
	public static void main(String[] args) {
		
		// create 2 boolean primitives bool1, bool2
		boolean bool1, bool2;
		
		// assign values to boo1, bool2
		bool1 = true;
		bool2 = false;
		
		// create 2 String's s1, s2
		String s1, s2;
		
		/**
		 *  static metyhod is called using class name
		 *  assign string values of priomitives bool1, bool2 to s1, s2
		 */
		s1 = Boolean.toString(bool1);
		s2 = Boolean.toString(bool2);
		
		String str1 = "String value fo boolean primitive " +bool1+ " is " +s1;
		String str2 = "String value fo boolean primitive " +bool2+ " is " +s2;
		
		// print s1, s2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}
