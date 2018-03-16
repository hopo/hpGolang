package level01;

import java.util.Arrays;

public class StrangeSort {
	public static void main(String args[]) {
		String[] sarr1 = {"sun", "bed", "car"};
		int n1 = 1;
		String[] ex1 = strangesort(sarr1, n1); // [car, bed, sun]
		System.out.println(Arrays.toString(ex1));
	}

	public static String[] strangesort(String[] sarr, int n) {
        int lnth = sarr.length;
        char[] carr = {};
        char[] box = new char[lnth];
        int b = 0;
        for(String str: sarr) {
            carr = str.toCharArray();
            box[b++] = carr[n];
        }

        char tmp;
        String stmp = "";
        for(int i = 0; i < lnth-1; i++) {
            for(int j = 1; j < lnth; j++) {
                if(box[i] > box[j]) {
                    tmp = box[i];
                    stmp = sarr[i];
                    box[i] = box[j];
                    sarr[i] = sarr[j];
                    box[j] = tmp;
                    sarr[j] = stmp;
                }
            }
        }

		return sarr;
	}
}
