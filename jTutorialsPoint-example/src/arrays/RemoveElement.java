package arrays;

import java.util.ArrayList;

public class RemoveElement {
    public static void main(String args[]) {
        ArrayList<Integer> arr = new ArrayList<Integer>(5);

        arr.add(20);
        arr.add(15);
        arr.add(30);
        arr.add(45);

        System.out.println("Size of list: " + arr.size());
        for(Integer number: arr) {
            System.out.println("Number = " + number);
        }

        arr.remove(2); // remove index 2
        System.out.println("Now, Size of list: " + arr.size());
        for(Integer number: arr) {
            System.out.println("Number = " + number);
        }
    }
}
