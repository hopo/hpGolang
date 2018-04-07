package level02;

public class ProductMatrix {
    public static void main(String args[]) {
        int[][] mata = {{1, 2}, {2, 3}};
        int[][] matb = {{3, 4}, {5, 6}};
    	int[][] ex1 = pdMatrix(mata, matb);	// [[13, 16], [21, 26]]
    }

    static int[][] pdMatrix(int[][] ma, int[][] mb) {
        int[][] ret = {{}, {}};
        
        return ret;
    }
}

/*
func productMatrix(A, B [][]int) [][]int {
	if len(A[0]) != len(B) {
		fmt.Println(" *** Operating impossible")
		return nil
	}

	l := len(A[0])
	var x, y int
	var xsl, ysl []int
	var xslsl, yslsl [][]int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			x = A[i][0] * B[0][j]
			y = A[i][1] * B[1][j]
			xsl = append(xsl, x)
			ysl = append(ysl, y)
		}
		xslsl = append(xslsl, xsl)
		yslsl = append(yslsl, ysl)
		xsl, ysl = []int{}, []int{}
	}

	var sumsl []int
	var ret [][]int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			sum := xslsl[i][j] + yslsl[i][j]
			sumsl = append(sumsl, sum)
		}
		ret = append(ret, sumsl)
		sumsl = []int{}
	}

	return ret
}
*/
