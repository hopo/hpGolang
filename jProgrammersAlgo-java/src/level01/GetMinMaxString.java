package level01;

import java.util.Arrays;

public class GetMinMaxString {
    public static void main(String args[]) {
        String str1 = "1 2 3 4";
        String[] ex1 = getminmax(str1); // [1, 4]
        System.out.println(Arrays.toString(ex1));
    }

    static String[] getminmax(String str) {
        String[] smpl = str.split("\\s");
        Arrays.sort(smpl);

        String[] ret = {smpl[0], smpl[smpl.length-1]};
        return ret;
    }
}
