// ...ing
package level02;

public class GetMinMax {
	public static void main(String[] args) {
		int[] arra1 = {1, 2}, arrb1 = {3, 4};
		int ex1 = getmm(arra1, arrb1); // 10

		int[] arra2 = {4, 2}, arrb2 = {3, 9};
		int ex2 = getmm(arra2, arrb2); // 30

		int[] arra3 = {2, 3}, arrb3 = {5, 7};
		int ex3 = getmm(arra3, arrb3); // 29
	}
	
	static int getmm(int[] arra, int[] arrb) {
		return -1;
	}
}
/*
func main() {
	ex1 := getMinMax([]int{1, 2}, []int{3, 4}) // 10
	ex2 := getMinMax([]int{4, 2}, []int{3, 9}) // 30
	ex3 := getMinMax([]int{2, 3}, []int{5, 7}) // 29
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)

}

func getMinMax(a, b []int) int {
	var data []int
	for _, v := range a {
		for _, w := range b {
			data = append(data, v*w)
		}
	}

	l := len(data)
	min := data[0] + data[l-1]
	for i := 0; i < l/2; i++ {
		if min > data[i]+data[l-1-i] {
			min = data[i] + data[l-1-i]
		}
	}
	return min
}
*/
