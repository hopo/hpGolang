package arraytest;

import java.util.ArrayList;

public class IntArrayAdd {
    public static void main(String args[]) {
        ArrayList<Integer> ilist = new ArrayList<>();

        Integer in1 = 42;
        Integer in2 = 17;
        Integer in3 = 8;
        int in4 = 7;

        ilist.add(in1);
        ilist.add(in2);
        ilist.add(in3);
        ilist.add(in4);

        // literal type int
        // and object Integer
        int total = 0;
        Integer ointTotal = 0;
       
        for(Integer oint : ilist) { total += oint; }
        for(int i : ilist) { ointTotal += i; }

        String sfmt = String.format("total: %d\n", total);
        System.out.println(sfmt);
        System.out.println("ointTotal: " + ointTotal);
    }
}
