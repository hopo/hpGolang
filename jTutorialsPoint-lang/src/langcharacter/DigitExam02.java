package langcharacter;

public class DigitExam02 {
	public static void main(String[] args) {
		// create 2 int primitives cp1, cp2
		int cp1, cp2;
		
		// assign values to cp1, cp2
		cp1 = 0x0034;
		cp2 = 0x004a;
		
		// create 2 int primitives i1, i2
		int i1, i2;
		
		// assign numeric value of cp1, cp2 to i1, i2 using radix
		i1 = Character.digit(cp1, 8);
		i2 = Character.digit(cp2, 8);
		
		String str1 ="Numeric value of cp1 in radix 8 is " + i1;
		String str2 ="Numeric value of cp2 in radix 8 is " + i2;
		
		// print i1, i2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}
