package level02;

public class IsPair {
	public static void main(String[] args) {
		String str1 = "(hello)()";
		boolean ex1 = pairChecker(str1); // true
		System.out.println(ex1);

		String str2 = ")(";
		boolean ex2 = pairChecker(str2); // false
		System.out.println(ex2);

		String str3 = "(5+(40*(1+2)/2)+(1+2)*2)";
		boolean ex3 = pairChecker(str3); // true
		System.out.println(ex3);
	}
	
	static boolean pairChecker(String str) {
		char[] chaArr = str.toCharArray();
		
		// is first open? is last close?
		boolean boolChk = (chaArr[0] == '(' && chaArr[chaArr.length-1] == ')');

		// is pair open and close?
		if(boolChk) {
			int op, cl;
			op = cl = 0;
			
			for(char c: chaArr) {
				switch (c) {
					case '(' :
						op++;
						break;
					case ')' :
						cl++;
						break;
				}
			}
			boolChk = (op == cl);
		}
		
		return boolChk;
	}

}
