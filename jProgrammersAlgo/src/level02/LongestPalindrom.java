package level02;

public class LongestPalindrom {
	public static void main(String[] args) {
		String str1 = "tomotmomtomot";
		int ex1 = lpalindrom(str1); // 13
		System.out.println(ex1);

		String str2 = "tomotmkdir";
		int ex2 = lpalindrom(str2); // 5
		System.out.println(ex2);
		
		String str3 = "xmmxaj";
		int ex3 = lpalindrom(str3); // 4
		System.out.println(ex3);
	}
	
	static int lpalindrom(String str) {
		char[] carr = str.toCharArray();

		int endPoint = carr.length;
		int endIdx = findEndIndex(carr, endPoint);
		int mChk, ret = 0;
		
		while(endIdx != 0) {
			mChk = mirrorCheck(carr, endIdx);
			if(endIdx != mChk) {
				endPoint = mChk;
				endIdx = findEndIndex(carr, endPoint);
			}
			else if(endIdx == mChk) {
				ret = endIdx;
				break;
			}
		}

		return ret+1;
		
	}
	
	static int findEndIndex(char[] carr, int point) {
		int ei = 0;

		for(int i = point-1; i != -1; i--) {
			if(carr[0] == carr[i]) {
				ei = i;
				break;
			}
		}
		
		return ei;
	}
	
	static int mirrorCheck(char[] carr, int endIdx) {
		int j, ret = endIdx;
		int scope = endIdx/2;

		for(int i = 0; i < scope; i++) {
			j = endIdx-i;
			if(carr[i] != carr[j]) {
				ret = j;
				break;
			}
		}

		return ret;
	}
}
