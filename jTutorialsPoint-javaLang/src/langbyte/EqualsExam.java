package langbyte;

public class EqualsExam {
    public static void main(String args[]) {
        // create 2 Byte objects b1, b2
        Byte b1, b2;

        // create 2 boolean primitives bool1, bool2
        boolean bool1, bool2;

        // assign values  to b1, b2
        b1 = new Byte("100");
        b2 = new Byte("100");

        // compare b1 and assign resylt to bool1
        bool1 = b1.equals(b2); // equal, b2 is Byte

        /**
         *  compare b1 with object 100 and assifn result to bool2, it
         *  returns false as 100 is not a Byte object
         */
        bool2 = b1.equals("100"); // not equal, "100" is String object

        String str1 = b1 + " equals " + b2 + " is " + bool1;
        String str2 = b1 + " equals object value 100 is " + bool2;

        // print bool1, bool2 values
        System.out.println(str1);
        System.out.println(str2);
    }
}
