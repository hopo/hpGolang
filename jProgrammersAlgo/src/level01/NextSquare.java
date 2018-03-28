package level01;

public class NextSquare {
    public static void main(String args[]) {
        int n1 = 121;
        int ex1 = next_square(n1); // 144, 12*12
        System.out.println(ex1);

        int n2 = 4;
        int ex2 = next_square(n2); // 9, 3*3
        System.out.println(ex2);

        int n3 = 5;
        int ex3 = next_square(n3); // 0
        System.out.println(ex3);
    }

    static int next_square(int n) {
        double ret = 0;
        double x = Math.sqrt(n);
        if(x == (int)x) {
            ret = Math.pow(x+1, 2);
        }

        return (int)ret;
    }
}
