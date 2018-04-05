package level02;

public class IsPair {
	public static void main(String[] args) {
		String str1 = "(hello)()";
		boolean ex1 = pairChecker(str1); // true
//		System.out.println(ex1);

		String str2 = ")(";
		boolean ex2 = pairChecker(str2); // false
//		System.out.println(ex2);

		String str3 = "(5+(40*(1+2)/2)+(1+2)*2)";
		boolean ex3 = pairChecker(str3); // true
//		System.out.println(ex3);
	}
	
	static boolean pairChecker(String str) {
		char[] chaArr = str.toCharArray();
		char[] smpl = new char[chaArr.length];
		int s = 0;
		
		for(char c: chaArr) {
			if(c == '(' || c == ')') {
				smpl[s++] = c;
			}
		}
		
		for(char c: smpl) {
			if(c == 0) { break; }
			System.out.print(c);
		}
		System.out.println();
		
		boolean ret = false;
		return ret;
	}

}
