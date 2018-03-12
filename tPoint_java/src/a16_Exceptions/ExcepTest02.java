package a16_Exceptions;

public class ExcepTest02 {

	public static void main(String[] args) {
		int a[] = new int[2];
		try {
			System.out.println("Accees element three :" + a[3]);
		} catch(ArrayIndexOutOfBoundsException e) {
			System.out.println("Excption thrown :" + e);
		} finally {
			a[0] = 6;
			System.out.println("First element value: " + a[0]);
			System.out.println("The finally statement is executed");
		}
	}
}
