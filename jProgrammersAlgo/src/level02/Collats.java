package level02;

public class Collats {
	public static void main(String[] args) {
		int num1 = 6;
		int ex1 = collats(num1); // 8, 3-10-5-16-8-4-2-1
		System.out.println(ex1);

		int num2 = 11;
		int ex2 = collats(num2); // 14, 34-17-52-26-13-40-20-10-5-16-8-4-2-1 
		System.out.println(ex2);
	}
	
	static int collats(int num) {
		int ret = 0;
		while(num != 1) {
			if(num%2 == 0) { num /= 2; }
			else { num = num*3+1; }
			ret++;
		}

		return ret;
	}
}