package level03;

public class Caesar {
	public static void main(String[] args) {
		String str1 = "a B z";
		int num1 = 4;
		String ex1 = csar(str1, num1); // "e F d"
		System.out.println(ex1);
		
		String str2 = "A b Z g";
		int num2 = 5;
		String ex2 = csar(str2, num2); //"F g E l"
		System.out.println(ex2);
	}
	
	static String csar(String str, int num) {
		char[] chArr = str.toCharArray();
		char c;
		for(int i = 0; i < chArr.length; i++) {
			c = chArr[i];
			if(c != ' ') {
				if(c == 90 || c == 122) {
					c -= 26;
				}
				chArr[i] = (char)(c+num);
			}
		}

		return new String(chArr);
	}

}