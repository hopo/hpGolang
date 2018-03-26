package langbyte;

// convert double type
public class DoubleValueExam {
    public static void main(String args[]) {
        // create 2 Byte obkects b1. b2
        Byte b1, b2;

        // create 2 double primitives d1, d2
        double d1, d2;

        // assign values to b1, b2
        b1 = new Byte("100");
        b2 = new Byte("-10");
            
        // assign values to b1, b2
        d1 = b1.doubleValue();
        d2 = b2.doubleValue();

        String str1 = "double value of Byte " + b1 + " is " + d1;
        String str2 = "double value of Byte " + b2 + " is " + d2;

        // print d1, d2 values
        System.out.println(str1);
        System.out.println(str1);
    }
}
