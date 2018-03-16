package level01;

import java.lang.String;

public class NoContinous {
    public static void main(String args[]) {
        String str1 = "13303";
        String ex1 = nocontinuous(str1); // "1 3 0 3"
        System.out.println(ex1);
    }

    public static String nocontinuous(String str) {
        char[] carr = str.toCharArray();
        String ret = "";

        for(int i = 1; i < carr.length; i++) {
            if(carr[i-1] != carr[i]) { ret += " "+ carr[i]; }
        }
        ret = carr[0] + ret;

        return ret;
    }
}
