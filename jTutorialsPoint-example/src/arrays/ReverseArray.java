package arrays;

public class ReverseArray {
    public static void main(String args[]) {
        int[] numbers = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

        System.out.println("Array before reverse:");
        // for(int i = 0; i < numbers.length; i++) {
        //     System.out.print(numbers[i] + " ");
        // }
        iarrPrint(numbers);

        // reversing >>
        for(int i = 0; i < numbers.length/2; i++) {
            int temp = numbers[i];
            numbers[i] = numbers[numbers.length-1-i];
            numbers[numbers.length-1-i] = temp;
        }
    
        System.out.println("\nArray after reverse:");
        // for(int i = 0; i < numbers.length; i++) {
            // System.out.print(numbers[i] + " ");
        // }
        iarrPrint(numbers);
    }

    static void iarrPrint(int[] iarr) {
        for(int i = 0; i < iarr.length; i++) {
            System.out.print(iarr[i] + " ");
        }
    }
}
