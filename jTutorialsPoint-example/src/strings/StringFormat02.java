package strings;

public class StringFormat02 {
    public static void main(String args[]) {
        String name = "Hello World";
        String s1 = String.format("name %s", name);
        String s2 = String.format("value %f", 32.33434);
        String s3 = String.format("value %32.12f", 32.33434);

        System.out.print(s1);
        System.out.print("\n");
        System.out.print(s2);
        System.out.print("\n");
        System.out.print(s3);
        System.out.print("\n");
    }
}
