package a13_RegularExpress;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class RegexMatches04 {
	
	private static String REGEX = "dog";
	private static String INPUT = "The dog says meow. " + "All dog say meow.";
	private static String REPLACE = "cat";

	public static void main(String[] args) {
		Pattern p = Pattern.compile(REGEX);
		
		// get a matcher object
		Matcher m = p.matcher(INPUT);
		INPUT = m.replaceAll(REPLACE);
		System.out.println(INPUT);
	}

}
