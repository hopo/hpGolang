package booleanEx;

import java.lang.*;

// static Boolean valueOf(boolean b)
public class BooleanDemo09 {
	public static void main(String[] args) {
		
		// create 2 Boolean objects b1, b2
		Boolean b1, b2;
		
		// create 2 boolean primitives and assign values
		boolean bool1 = true;
		boolean bool2 = false;
		
		/**
		 *  static method is called using class name
		 *  assign result of valueOf method on bool1, bool2 to b1, b2
		 */
		b1 = Boolean.valueOf(bool1);
		b2 = Boolean.valueOf(bool2);
		
		String str1 = "Boolean instance of primitive " +bool1 + " is " + b1;
		String str2 = "Boolean instance of primitive " +bool2 + " is " + b2;
		
		// print b1, b2 values
		System.out.println(str1);
		System.out.println(str2);
	}

}
