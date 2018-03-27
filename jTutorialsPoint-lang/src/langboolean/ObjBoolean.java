package langboolean;

// boolean equals(Object obj)
public class ObjBoolean {
    public static void main(String args[]) {

        // create 2 Boolean objects b1, b2
        Boolean b1, b2;

        // create a boolean  primitive res
        boolean res;

        // assign values to b1, b2
        b1 = new Boolean(true);
        b2 = new Boolean(false);

        // assign the result of equals method on b1, b2 to res
        res = b1.equals(b2);

        String str = "b1:" +b1+ " and b2:" +b2+ " are equal is " +res;

        // print res value
        System.out.println(str);
    }
}