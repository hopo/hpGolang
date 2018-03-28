package level01;

import java.util.Arrays;

public class RmSmall {
    public static void main(String args[]) {
        int[] iarr1 = new int[]{4, 3, 2, 1};
        int[] ex1 = rmsmall(iarr1); // [4, 3, 2]
        System.out.println(Arrays.toString(ex1));

        int[] iarr2 = new int[]{10, 8, 22};
        int[] ex2 = rmsmall(iarr2); // [10, 22]
        System.out.println(Arrays.toString(ex2));
    }

    static int[] rmsmall(int[] iarr) {
        int r = 0, sm = iarr[0];
        int[] ret = new int[iarr.length-1];

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
