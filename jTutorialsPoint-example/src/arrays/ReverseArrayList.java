package arrays;

import java.util.ArrayList;
import java.util.Collections;

public class ReverseArrayList {
    public static void main(String args[]) {
        ArrayList<String> arrayList = new ArrayList<String>();
        arrayList.add("A");
        arrayList.add("B");
        arrayList.add("C");
        arrayList.add("D");
        arrayList.add("E");

        System.out.println("Before Reverse Order: " + arrayList);

        Collections.reverse(arrayList);
        System.out.println("After Reverse Order: " + arrayList);
    }
}
