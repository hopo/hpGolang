package langboolean;

// boolean booleanValue()
public class BooleanExam {
    public static void main(String args[]) {

        // create a Boolean object b
        Boolean b;
        
        // assign value to b
        b = new Boolean(true);

        // create a boolean primitive type bool
        boolean bool;

        // assign primitive value of b to bool
        bool = b.booleanValue();

        String str = "Primitive value of Boolean object " + b + " is "  + bool;

        // print boool value
        System.out.println(str);
    }
}
