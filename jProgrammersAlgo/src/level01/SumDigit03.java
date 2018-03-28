package level01;

public class SumDigit03 {
    public static void main(String args[]) {
        int num1 = 123;
        int ex1 = sumdigit(num1); // 6
        System.out.println(ex1);

        int num2 = 999999999;
        int ex2 = sumdigit(num2); // 81
        System.out.println(ex2);
    }

    static int sumdigit(int num) {
        int ret = 0;
        int[] iarr = eachIntArr(num);
        for(int i : iarr) { ret += i; }

        return ret;
    }
    
    static int[] eachIntArr(int n) {
        int[] ret = new int[64];
        int z, r = 0, di = 1;

        while(true) {
            if(n-di < 0){
                di /= 10;
                break;
            }
            di *= 10;
        }

        while(true) {
            if(di == 1){
                ret[r++] = n;
                break;
            }
            z = n/di;
            ret[r++] = z;
            n %= di;
            di /=  10;
        }

        return ret;
    }
}
