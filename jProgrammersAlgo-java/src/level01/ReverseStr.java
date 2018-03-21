package level01;

import java.util.Arrays;

public class ReverseStr {
	public String reverseStr(String str){
    		String ret = "";
        char[] ctmp = str.toCharArray();
        Arrays.sort(ctmp);
    		for(int i = 0; i < ctmp.length; i++) {
    			ret = Character.toString(ctmp[i]) + ret;
        }

    return ret;
	}

	// 아래는 테스트로 출력해 보기 위한 코드입니다.
	public static void main(String[] args) {
		ReverseStr rs = new ReverseStr();
		System.out.println( rs.reverseStr("Zbcdefg") );
	}
}
