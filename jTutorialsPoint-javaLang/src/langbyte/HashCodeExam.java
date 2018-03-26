package langbyte;

// Byte value convert to int value
public class HashCodeExam {
    public static void main(String args[]) {
        // create 2 Byte objects b1, b2
        Byte b1, b2;

        // create 2 int i1, i2
        int i1, i2;

        // assifn values to b1, b2
        b1 = new Byte("100");
        b2 = new Byte("-100");

        // assign hash code of b1, b2 to i1, i2
        i1 = b1.hashCode();
        i2 = b2.hashCode();

        String str1 = "Hashcode of Byte " + b1 + " is " + i1;
        String str2 = "Hashcode of Byte " + b2 + " is " + i2;

        // print i1, i2 values
        System.out.println(str1);
        System.out.println(str2);
    }
}
