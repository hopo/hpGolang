package level02;

public class LongestPalindrom99 {
	public static void main(String[] args) {
		String str1 = "tomotmomtomot";
		int ex1 = lpalindrom(str1); // 13
		System.out.println(ex1);

		/*
		String str2 = "tomotmkdir";
		int ex2 = lpalindrom(str2); // 5
		System.out.println(ex2);
		
		String str3 = "xmmxaj";
		int ex3 = lpalindrom(str3); // 4
		System.out.println(ex3);
		*/
	}
	
	static int lpalindrom(String str) {
		char[] carr = str.toCharArray();
		char start = carr[0];
		char[] sample = sameSE(start, carr);
		
		int ret = -1;
		if(mirrorCheck(sample)) {
			ret = sample.length;
		}
		
		return ret;
	}
	
	static char[] sameSE(char startchar, char[] smplcarr) {
		int r = 0;
		int lnth = smplcarr.length;
		for(int i = lnth-1; i > 0; i--) {
			if(smplcarr[i] == startchar) {
				r = i;
				break;
			}
		}
		
		char[] ret = new char[r+1];
		for(int i = 0; i < r+1; i++) {
			ret[i] = smplcarr[i];
		}

		return ret;
	}
	
	static boolean mirrorCheck(char[] arrangedcarr) {
		int lnth = arrangedcarr.length;
		char[] fhalf = new char[lnth/2+1];
		char[] shalf;
		if(lnth%2 == 0) {
			shalf = new char[lnth/2+1];
		} else {
			shalf = new char[lnth/2];
		}
		
		for(int i = 0; i < fhalf.length-1; i++) {
			fhalf[i] = arrangedcarr[i];
		}
		
		for(int i = 0; i < shalf.length; i++) {
			if(lnth%2 == 0) {
				shalf[i] = arrangedcarr[i+lnth/2+1];
			} else {
				shalf[i] = arrangedcarr[i+lnth/2];
			}
		}
	
		boolean ret = true;
		/*
		int j;
		for(int i = 0; i < fhalf.length; i++) {
			j = lnth/2 - 1 - i;
			if(fhalf[i] != shalf[j]) {
				ret = false;
				break;
			}
		}
		*/

		return ret;
	}
}
