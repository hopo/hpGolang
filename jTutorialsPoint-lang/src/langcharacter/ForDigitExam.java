package langcharacter;

public class ForDigitExam {
	public static void main(String[] args) {
		// create 2 character primitives ch1, ch2
		char ch1, ch2;
		
		// create 2 int primitives i1, i2 and assign values
		int i1 = 3;
		int i2 = 14;
		
		// assign char representation of i1, i2 to ch1, ch2 using radix
		ch1 = Character.forDigit(i1, 10);
		ch2 = Character.forDigit(i2, 16);
		
		String str1 = "Char representation of " + i1 + " in radix 10 is " + ch1;
		String str2 = "Char representation of " + i2 + " in radix 10 is " + ch2;
		
		// print ch1, ch2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}