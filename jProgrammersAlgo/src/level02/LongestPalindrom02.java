package level02;

public class LongestPalindrom02 {
	public static void main(String[] args) {
		/*
		String str1 = "tomotmomtomot";
		int ex1 = lpalindrom(str1); // 13
		System.out.println(ex1);

		String str2 = "tomotmkdir";
		int ex2 = lpalindrom(str2); // 5
		System.out.println(ex2);
		
		String str3 = "xmmxaj";
		int ex3 = lpalindrom(str3); // 4
		System.out.println(ex3);
		*/

		String str4 = "xxvx22";
		int ex4 = lpalindrom(str4);
		System.out.println(ex4);
	}
	
	static int lpalindrom(String str) {
		char[] charArr = str.toCharArray();
		int endPoint = charArr.length;
		
		int lastIdx = findLIdx(charArr, endPoint);
		int mChk = mirrorable(charArr, lastIdx);
		
		System.out.println(lastIdx);
		System.out.println(mChk);

		return -1;
		
	}
	
	static int findLIdx(char[] charArr, int endPoint) {
		char start = charArr[0];
		int lastIdx = 0;
		
		for(int i = charArr.length-1; i > -1; i++) {
			if(start == charArr[i]) {
				lastIdx = i;
				break;
			}
		}
		
		return lastIdx;
	}
	
	static int mirrorable(char[] carr, int idx) {
		int j, ret = idx;
		int scope = idx/2;

		for(int i = 0; i < scope; i++) {
			j = idx-i;
			if(carr[i] != carr[j]) {
				ret = j;
				break;
			}
		}

		return ret;
	}
}
