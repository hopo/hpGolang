package arrays;

import java.util.Arrays;

public class FillTest {
    public static void main(String args[]) {
        int[] array = new int[6];
        Arrays.fill(array, 100); // Fill 100 array each element
        for(int i = 0, n = array.length; i < n; i++) {
            System.out.print(array[i]);
            System.out.print(" ");
        }
        System.out.println();

        Arrays.fill(array, 3, 6, 50); // Fill 50 array[3] ~ array[5]
        for(int i = 0, n = array.length; i < n; i++) {
            System.out.print(array[i]);
            System.out.print(" ");
        }
    }
}
