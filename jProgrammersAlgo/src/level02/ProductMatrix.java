package level02;

import java.util.Arrays;

public class ProductMatrix {
    public static void main(String args[]) {
    	// 2x2
		int[][] mata1 = {{1, 2}, {2, 3}};
		int[][] matb1 = {{3, 4}, {5, 6}};
    	int[][] ex1 = pdMatrix(mata1, matb1);	// [[13, 16], [21, 26]]
    	String ppp = Arrays.deepToString(ex1);
    	System.out.println(ppp);

    	// 3x3
        int[][] mata2 = {{1, 2, 2}, {2, 3, 4}, {3, 3, 4}};
        int[][] matb2 = {{3, 4, 7}, {5, 6, 8}, {2, 4, 1}};
    	int[][] ex2 = pdMatrix(mata2, matb2);	// [[17, 24, 25], [29, 42, 42], [32, 46, 49]]	

    	ppp = Arrays.deepToString(ex2);
    	System.out.println(ppp);
    }

    static int[][] pdMatrix(int[][] ma, int[][] mb) {
    	int raw = ma.length;
    	int col = ma[0].length;
    	int ele = 0;
    	int[][] ret = new int[raw][col];
    	if(raw == col) { 
			for(int i = 0; i < raw; i++) {
				for(int j = 0; j < col; j++) {
					for(int k = 0; k < raw; k++) {
						ele += ma[i][k] * mb[k][j];
					}
					ret[i][j] = ele;
					ele = 0;
				}
			}
    	}

        return ret;
    }
}
