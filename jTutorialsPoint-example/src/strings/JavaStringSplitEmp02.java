package strings;

public class JavaStringSplitEmp02 {
    public static void main(String args[]) {
        String s1 = "t u t o r i 1 l s";
        String[] words = s1.split("\\s");
        for(String w : words) {
            System.out.println(w);
        }
    }
}
