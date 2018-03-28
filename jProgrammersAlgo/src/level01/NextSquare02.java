package level01;

public class NextSquare02 {
    public static void main(String args[]) {
        int n1 = 121;
        int ex1 = next_square(n1); // 144, 12*12
        System.out.println(ex1);

        /*
        int n2 = 4;
        int ex2 = next_square(n2); // 9, 3*3
        System.out.println(ex2);

        int n3 = 5;
        int ex3 = next_square(n3); // 0
        System.out.println(ex3);
        */
    }

    static int next_square(int n) {
        double ret = 0;
        double x = Math.sqrt(n);
        if(x == (int)x) {
            ret = Math.pow(x+1, 2);
        }

        return (int)ret;
    }
}

/*
func main() {
	ex1 := nextSqaure(121) // 144, 12*12
	ex2 := nextSqaure(4)   // 9, 3*3
	ex3 := nextSqaure(5)   // 0, "no"
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
}

func nextSqaure(n int) (rslt int) {
	x := math.Sqrt(float64(n))
	f := math.Pow(x, 2)
	if f == float64(n) {
		rslt = int(math.Pow(x+1, 2))
	} else {
		rslt = -1
	}
	return
}
	"math"
)
*/

