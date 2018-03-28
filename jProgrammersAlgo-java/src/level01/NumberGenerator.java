package level01;

import java.util.Arrays;

public class NumberGenerator {
    public static void main(String args[]) {
        int x1 = 3, n1 = 5;
        int[] ex1 = numgenerator(x1, n1); // [3, 6, 9, 12, 15]
        System.out.println(Arrays.toString(ex1));

        int x2 = 4, n2 = 3;
        int[] ex2 = numgenerator(x2, n2); // [4, 8, 12]
        System.out.println(Arrays.toString(ex2));
    }

    static int[] numgenerator(int x, int n) {
        int r = 0;
        int[] ret = new int[n];

        for(int i = 1; i < n+1; i++) {
            ret[r++] = x*i;
        }

        return ret;
    }
}
