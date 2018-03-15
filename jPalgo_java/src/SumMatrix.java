// package level01;

import java.util.Arrays;

public class SumMatrix {

    public static void main(String args[]) {
        SumMatrix sm1 = new SumMatrix();
        int[][] a1 = {{1, 2}, {2, 3}}; 
        int[][] b1 = {{3, 4}, {5, 6}}; 

        int[][] ex1 = sm1.sumMatrix(a1, b1); // [[4, 6], [7, 9]]
        System.out.print("[" + Arrays.toString(ex1[0]));
        System.out.println(Arrays.toString(ex1[1]) + "]");
    }

    public int[][] sumMatrix(int[][] a, int[][] b) {
        int ret[][] = {{0, 0}, {0, 0}};
        for(int i = 0; i < 2; i++)
            for(int j = 0; j < 2; j++)
                ret[i][j] = a[i][j] + b[i][j];

        return ret;
    }
}
