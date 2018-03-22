package level01;

public class SumDigit {
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
        Integer oint = new Integer(num);
        String ostr = oint.toString();
        char[] carr = ostr.toCharArray();

        for(char c : carr) {
            ret += Character.getNumericValue(c);
        }

        return ret;
    }
}
