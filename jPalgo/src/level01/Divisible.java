// ing
// package level01;

import java.util.Arrays;

public class Divisible {
    public static void main(String args[]) {
        int[] iarr1 = {5, 9, 7, 10};
        int n1 = 5;
        Divisible di1 = new Divisible();
        int[] ex1 = di1.divisible(iarr1, n1); // [5, 10]
        System.out.println(Arrays.toString(ex1));
    }

    int[] divisible(int[] iarr, int n) {
        int[] ret = {};
        for(int i : iarr)
            if(i%n == 0)
                System.out.println(i);
        return ret;
    }
}


/*
func main() {
	ex := divisible([]int{5, 9, 7, 10}, 5) // [5 10]
	fmt.Println(ex)
}

func divisible(isl []int, in int) (ret []int) {
	for _, v := range isl {
		if v%in == 0 {
			ret = append(ret, v)
		}
	}
	return
}
*/
