package tutorial.filesandio;

import java.io.File;

public class Main01 {

	public static void main(String[] args) {
		File file1 = new File("./demo1.txt");
		File file2 = new File("./demo2.txt");
		
		if(file1.compareTo(file2) == 0) {
			System.out.println("Both path are same!");
		} else {
			System.out.println("Path are not same!");
		}
	}
}