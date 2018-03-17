package booleanEx;

import java.lang.*;

// static Boolean valueOf(String s)
public class BooleanDemo10 {
	public static void main(String[] args) {
		
		// create 2 Boolean objects b1, b2
		Boolean b1, b2;
		
		// create 2 String's and assign values
		String s1 = null;
		String s2 = "false";
		
		/**
		 *  staitc method si called using name
		 *  assign result of valueOf method on s1, s2 to b1, b2
		 */
		b1 = Boolean.valueOf(s1);
		b2 = Boolean.valueOf(s2);
		
		String str1 = "Boolean instance of string " +s1+ " is " +b1;
		String str2 = "Boolean instance of string " +s2+ " is " +b2;
		
		// prtint b1, b2 values
		System.out.println(str1);
		System.out.println(str2);
	}

}
