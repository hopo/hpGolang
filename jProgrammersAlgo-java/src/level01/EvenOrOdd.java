package level01;

public class EvenOrOdd {
    public static void main(String args[]) {
        int num1 = 3;
        String ex1 = evenorodd(num1); // "Odd"
        System.out.println(ex1);

        int num2 = 2;
        String ex2 = evenorodd(num2); // "Even"
        System.out.println(ex2);
    }

    static String  evenorodd(int num) {
        return (num%2 == 0)? "Even" : "Odd";
    }
}
