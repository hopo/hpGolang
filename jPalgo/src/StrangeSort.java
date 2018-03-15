// ing
// package level01;

import java.util.Arrays;

public class StrangeSort {
	public static void main(String args[]) {
		StrangeSort ss1 = new StrangeSort();
		String[] sarr1 = {"sun", "bed", "car"};
		int n1 = 1;
		String[] ex1 = ss1.ssorts(sarr1, n1); // [car, bed, sun]
		System.out.println(Arrays.toString(ex1));
	}

	String[] ssorts(String[] sarr, int n) {
		String[] ret = {};
		return ret;
	}
}

/*
func main() {
	ex := strange_sort([]string{"sun", "bed", "car"}, 1) // [car bed sun]
	fmt.Println(ex)
}

const ALP = "abcdefghijklmnopqrstuvwxyz"

func strange_sort(ssl []string, n int) (ret []string) {
	for _, r := range ALP {
		for _, v := range ssl {
			if byte(r) == v[n] {
				ret = append(ret, v)
			}
		}
	}
	return
}
*/
