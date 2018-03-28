package stringtest;

public class StringBufferExam {
    public static void main(String args[]) {
        String str = "Hello world";

        StringBuffer sb = new StringBuffer(str);

        sb.append("!");
        sb.insert(5, " -");
        System.out.println(sb);
    }
}
