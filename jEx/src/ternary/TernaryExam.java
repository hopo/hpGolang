package ternary;

public class TernaryExam {
    public static void main(String args[]) {
        int smp, ch;
        boolean bbb;

        for(int i = 0; i < 10; i++) {
            smp = (int)(Math.random()*100 + 1);
            ch  = smp&1;
            bbb = (ch == 1)? true : false;
            System.out.printf("%3d is Odd? %b\n", smp, bbb);
        }
    }
}
