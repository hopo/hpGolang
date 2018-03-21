package level01;

import java.util.Arrays;

public class Divisible {
    public static void main(String args[]) {
        int[] iarr1 = {5, 9, 7, 10};
        int n1 = 5;
        int[] ex1 = divisible(iarr1, n1); // [5, 10]
        System.out.println(Arrays.toString(ex1));
    }

    public static int[] divisible(int[] iarr, int n) {
        int b = 0, lnth = iarr.length;
        int[] box = new int[lnth];

        for(int i : iarr) {
            if(i%n == 0) {
                box[b++] = i;
            }
        }

        int[] ret = new int[b];
        for(int i = 0; i < b; i++) { ret[i] = box[i]; }

        return ret;
    }
}
