package langbyte;

public class CompareToExam {
    public static void main(String args[]) {
        // creatw 2 Byte objects b1, b2
        Byte b1, b2;

        // assign values to b1, b2
        b1 = new Byte("-100");
        b2 = new Byte("10");

        // create an in res
        int res;

        // compare b1 with b2 and assign resut to res
        res = b1.compareTo(b2); // return 10

        String str1 = "Both values are equal ";
        String str2 = "First value is greater";
        String str3 = "Second value is greater";

        if(res == 0) {
            System.out.println(str1);
        }
        else if(res > 0) {
            System.out.println(str2);
        }
        else if(res < 0) {
            System.out.println(str3);
        }
    }
}
