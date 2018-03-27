package arrays;

import java.util.Arrays;

public class ArraysDeeptostringExam {
	public static void main(String[] args) {
		String[][] deepArr = new String[][] {{"Sai", "Gopar"}, {"Pallavi", "Priya"}};
		System.out.println(Arrays.toString(deepArr));
		System.out.println(Arrays.deepToString(deepArr));
	}
}
