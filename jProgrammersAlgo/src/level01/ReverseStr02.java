package level01;

public class ReverseStr02 {
	public static void main(String[] args) {
        String str1 = "Zbcdefg";
        String ex1 = reversestr(str1); // "gfedcbZ"
		System.out.println(ex1); 
	}

	static String reversestr(String str){
        return new StringBuffer(str).reverse().toString();
	}
}
