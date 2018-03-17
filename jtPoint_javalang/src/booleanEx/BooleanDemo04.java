package booleanEx;

// static boolean getProperty(String name)
import java.lang.*;

public class BooleanDemo04 {
    public static void main(String args[]) {
        
        // create 2 boolean primitives bool1, bool2
        boolean bool1, bool2;

        /**
         *  using System class's setProprty method, set the values of
         *  system properties demo1, demo2.
         */
        System.setProperty("demo1", "true");
        System.setProperty("demo2", "abcd");

        // retrieve value of system properties to s1, s2
        String s1 = System.getProperty("demo1");
        String s2 = System.getProperty("demo2");

        // assign result of getBoolean on demo1, demo2 to bool1, bool2
        bool1 = Boolean.getBoolean("demo1");
        bool2 = Boolean.getBoolean("demo2");

        String str1 = "boolean value of system property demo1 is " + bool1;
        String str2 = "System property value of demo1 is " + s1;
        String str3 = "boolean value of system property demo2 is " + bool2;
        String str4 = "System property value of demo1 is " + s2;

        // print bool1, bool2 and s1, s2 values
        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);
        System.out.println(str4);
    }
}
