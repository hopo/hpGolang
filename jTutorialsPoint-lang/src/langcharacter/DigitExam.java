package langcharacter;

public class DigitExam {
	public static void main(String[] args) {
		// create 2 chatacter primitives ch1, ch2
		char ch1, ch2;
		
		// assign values to ch1, ch2
		ch1 = '9';
		ch2 = '5';
		
		// create 2 int primitives i1, u2
		int i1, i2;
		
		// assign numeric value of ch1, ch2 to i1, i2 using radix
		i1 = Character.digit(ch1, 2);
		i2 = Character.digit(ch2, 10);
		
		String str1 = "Numeric value of " + ch1 + " in radix 2 is " + i1;
		String str2 = "Numeric value of " + ch2 + " in radix 10 is " + i2;
		
		// print i1, i2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}
