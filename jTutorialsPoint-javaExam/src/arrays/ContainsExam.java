package arrays;

import java.util.ArrayList;

public class ContainsExam {
    public static void main(String args[]) {
        ArrayList<String> objArray = new ArrayList<>();
        ArrayList<String> objArray2 = new ArrayList<>();

        objArray.add(0, "common1");
        objArray.add(1, "common2");
        
        objArray2.add(0, "common1");
        objArray2.add(1, "common2");
        objArray2.add(2, "notcommon");
        objArray2.add(3, "notcommon1");

        System.out.println("Array elements of array1" + objArray);
        System.out.println("Array elements of array2" + objArray2);
        System.out.println(
            "Array 1 contains String common2?? "       
            + objArray.contains("common1")
        ); // true
        System.out.println(
            "Array 2 contains Array1?? "
            + objArray2.contains(objArray)
        ); // false
    }
}
