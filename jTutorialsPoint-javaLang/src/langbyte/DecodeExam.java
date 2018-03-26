package langbyte;

public class DecodeExam {
    public static void main(String args[]) {
        // create 4 Byte objects
        Byte b1, b2, b3, b4;

        /*
         *  static methods are called using class name.
         *  decimal value is decoded and assigned to Byte objects b1
         */
        b1 = Byte.decode("100");

        // hexadecimal values are decoded and assigned to Byte objects
        b2 = Byte.decode("0x6b");
        b3 = Byte.decode("-#4c");

        // octal values is decoded and assigned to Byte object b4
        b4 = Byte.decode("0127");

        String str1 = "Byte value of decimal 100 is " + b1;
        String str2 = "Byte value of hexadecimal 6b is " + b2;
        String str3 = "Byte value of hexadecimal -4c is " + b3;
        String str4 = "Byte value of octal 127 is " + b4;

        // print b1, b2, b3, b4 values
        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);
        System.out.println(str4);

    }
}
