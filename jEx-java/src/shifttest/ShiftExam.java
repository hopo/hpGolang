package shifttest;

public class ShiftExam {
    public static void main(String args[]) {
        int x = 42;

        System.out.printf("x: 0x%04x(%d)\n", x, x);

        // System.out.println();
        // System.out.printf("x/2: %d\n", x/2);
        // System.out.printf("x/2: %d\n", (int)(x*0.5));

        // rigth shift 1, 2, 3
        System.out.println();
        for(int i = 1; i < 4; i++) {
            System.out.printf("x>>%d: 0x%04x(%d): %.2f\n", i, x>>i, x>>i, x*1/Math.pow(2, i));
        }

        // left shift 1, 2, 3
        System.out.println();
        for(int i = 1; i < 4; i++) {
            System.out.printf("x<<%d: 0x%04x(%d): %.2f\n", i, x<<i, x<<i, x*Math.pow(2, i));
        }

        // make all Even number
        System.out.println();
        int t = 0;
        for(int i = 0; i < 20; i++) {
            t = i<<1;
            System.out.print(t&1); //if t is Odd then print 1
            System.out.print(" ");
        }
    }
}
