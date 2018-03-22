package strings;

public class StringBufferInsert {
    public static void main(String args[]) {
        StringBuffer sb = new StringBuffer("hello");
        sb.append(" world");
        sb.insert(0, "Tutorialspoint ");
        System.out.print(sb);
    }
}
