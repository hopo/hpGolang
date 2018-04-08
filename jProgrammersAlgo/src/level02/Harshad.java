package level02;

public class Harshad {
	public static void main(String[] args) {
		int iNum1 = 18;
		boolean ex1 = isHar(iNum1); // true, 1+8=9; 18%9 == 0
		System.out.println(ex1);

		int iNum2 = 13;
		boolean ex2 = isHar(iNum2); // false, 1+3=4; 13%4 != 0
		System.out.println(ex2);

		int iNum3 = 342;
		boolean ex3 = isHar(iNum3); // true, 3+4+2=9; 342%9 == 0
		System.out.println(ex3);
	}
	
	static boolean isHar(int num) {
		char[] chArr = Integer.toString(num).toCharArray();
		
		int sum = 0;
		for(char c: chArr) {
			sum += c-48;
		}
		
		return (num%sum == 0)? true : false;
	}

}
