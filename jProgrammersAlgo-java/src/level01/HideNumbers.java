package level01;

import java.util.Arrays;

public class HideNumbers {
    public static void main(String args[]) {
        String pnum1 = "01033334444";
        String ex1 = hidenumber(pnum1); // "*******4444"
        System.out.println(ex1);

        String pnum2 = "027778888";
        String ex2 = hidenumber(pnum2); // "*****8888"
        System.out.println(ex2);
    }

    static String hidenumber(String pnum) {
        char[] carr = pnum.toCharArray();
        String ret = "";

        for(int i = 0; i < carr.length; i++) {
            if(i < carr.length-4) { ret += "*"; }
            else { ret += carr[i]; }
        }

        return ret;
    }
}
