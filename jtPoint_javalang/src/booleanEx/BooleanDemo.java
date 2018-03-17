package booleanEx;

import java.lang.*;

// booleanValue()
public class BooleanDemo {
    public static void main(String args[]) {
        // create a Boolean object b
        Boolean b=true;
        
        // assign value to b (warning message)
        // b = new Boolean(true);

        // create a boolean primitive type bool
        boolean bool;

        // assign primitive value of b to bool
        bool = b.booleanValue();

        String str = "Primitive value of Boolean object " + b + " is "  + bool;

        // print boool value
        System.out.println(str);
    }
}
