// package level01;

public class StringMiddle {
	public static void main(String[] args) {
		StringMiddle sm1 = new StringMiddle();
		String str1 = new String("power");
		String ex1 = sm1.stringmiddle(str1); // w
		System.out.println(ex1);

/*
		StringMiddle sm2 = new StringMiddle();
		String str2 = new String("java");
		String ex2 = sm2.stringmiddle(str2); // av
		System.out.print(ex2);
*/
	}

	String stringmiddle(String str) {
		int str_len = str.length();
		int str_len_half = str_len/2;
		/*
		if(str_len%2 == 0) {
			return str[str_len_half-1] + len[str_len_half+1];
		}
		*/
		return "break! String";
	}
}


/*
return len(str)&1 and str[len(str)//2] or str[len(str)//2-1:len(str)//2+1]
*/
