package langcharacter;

public class CharValueExam {
	public static void main(String[] args) {
		// create a Character object c
		Character c;
		
		// assign value to c
		c = new Character('a');
		
		// create a char primitve ch
		char ch;
		
		// assign primitive value of c to ch
		ch = c.charValue();
		
		String str = "Primitive char value is " + ch;
		
		// print ch value
		System.out.println(str);
	}
}
