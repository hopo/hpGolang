package level02;

public class Adder {
	public static void main(String[] args) {
		int a1 = 3, b1 = 5;
		int ex1 = adder(a1, b1); // 12, 3+4+5;
		System.out.println(ex1);

	}
	
	static int adder(int a, int b) {
		int ret = b;
		
		// ret = a + ... + b-1
		for(int i = a; i < b; i++) { ret += i; }
		
		return ret;
	}
}
