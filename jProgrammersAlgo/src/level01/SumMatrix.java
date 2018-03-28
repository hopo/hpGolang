package level01;

import java.util.Arrays;

public class SumMatrix {

    public static void main(String args[]) {
        int[][] a1 = {{1, 2}, {2, 3}}; 
        int[][] b1 = {{3, 4}, {5, 6}}; 

        int[][] ex1 = summatrix(a1, b1); // [[4, 6], [7, 9]]
        System.out.print("[" + Arrays.toString(ex1[0]));
        System.out.println(Arrays.toString(ex1[1]) + "]");
    }

    public static int[][] summatrix(int[][] a, int[][] b) {
        int ret[][] = new int[a.length][a[0].length];
        for(int i = 0; i < 2; i++) {
            for(int j = 0; j < 2; j++) {
                ret[i][j] = a[i][j] + b[i][j];
            }
        }

        return ret;
    }
}
