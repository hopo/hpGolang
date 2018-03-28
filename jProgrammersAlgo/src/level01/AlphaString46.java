package level01;

import java.lang.String;

public class AlphaString46 {
    public static void main(String[] args) {
        String str1 = "a234";
	    boolean ex1 = alpha_string46(str1); // false, "string 'a'"
        System.out.println(ex1);

        String str2="1234";
	    boolean ex2 = alpha_string46(str2); // true, "all int"
        System.out.println(ex2);
    }

    static boolean alpha_string46(String str){
        boolean ret = true;
        char smpl;
        for(int i = 0; i < str.length(); i++) {
            smpl = str.charAt(i);
            if('0' > smpl || smpl > '9') {
                ret = false;
                break;
            }
        }
	    return ret;
    }
}
