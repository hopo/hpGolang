package level02;

public class ToWeirdCase {
	public static void main(String[] args) {
		String str1 = "try hello world";
		String ex1 = towcase(str1); // "TrY HeLlO WoRlD"
		System.out.println(ex1);
		
		String str2 = "coding is thinking not typing";
		String ex2 = towcase(str2); // "CoDiNg Is ThInKiNg NoT TyPiNg"
		System.out.println(ex2);
	}
	
	static String towcase(String str) {
		char[] chaArr = str.toCharArray();
		int flag = 0;
		
		for(int i = 0; i < chaArr.length; i++) {
			if(chaArr[i] == ' ') {
				flag = 0;
				continue;
			}

			flag++;
			if((flag&1) == 0) { continue; }
			else { chaArr[i] -= 32; }
		}
		
		return new String(chaArr);
	}

}
