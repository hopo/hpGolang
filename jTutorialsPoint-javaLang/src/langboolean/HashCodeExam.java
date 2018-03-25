package langboolean;

// int hashcode()
public class HashCodeExam {
	public static void main(String[] args) {

		// create 2 Boolean objects b1, b2
		Boolean b1, b2;
		
		// assing values to b1, b2
		b1 = true;
		b2 = false;
		
		// create 2 int primitives
		int i1, i2;
		
		// assign the hash code of b1, b2 to i1, i2
		i1 = b1.hashCode();
		i2 = b2.hashCode();
		
		String str1 = "Hash code of" +b1+ " is " +i1;
		String str2 = "Hash code of" +b2+ " is " +i2;
		
		// print i1, i2 values
		System.out.println(str1);
		System.out.println(str2);
	}
}
