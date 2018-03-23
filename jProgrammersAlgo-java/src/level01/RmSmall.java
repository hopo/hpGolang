package level01;

import java.util.Arrays;

public class RmSmall {
    public static void main(String args[]) {
        int[] iarr1 = {4, 3, 2, 1};
        int[] ex1 = RmSmall(iarr1); // [4 3 2]
        System.out.println(Arrays.toString(ex1));

        int[] iarr2 = {10, 8, 22};
        int[] ex2 = RmSmall(iarr2); // [10 22]
        System.out.println(Arrays.toString(ex2));
    }

    static int[] RmSmall(int[] iarr) {
        int[] ret = new int[iarr.length-1];
        int r = 0, sm = iarr[0];
        for(int i = 1; i < iarr.length; i++) {
            if(sm > iarr[i]) { sm = iarr[i]; }
        }

        for(int i = 0; i < iarr.length; i++) {
            if(iarr[i] == sm) { continue; }
            ret[r++] = iarr[i];
        }
        
        return ret;
    }
}
