package langbyte;

public class ByteDemo {
    public static void main(String args[]) {
        // create a Byte object b
        Byte b;

        // assign value to b
        b = new Byte("100");

        // create a byte primitive bt
        byte bt;

        // assign primitive value of b to bt
        bt = b.byteValue();

        String str = "Primitive byre value of Byte objecet " + b + " is " + bt;

        // print bt value
        System.out.println(str);
    }
}
