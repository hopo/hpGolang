package level01;

public class PrintTriangle {
    public static void main(String args[]) {
        int n1 = 3;
        String ex1 = print_triangle(n1);
        System.out.println(ex1);
    }

    public static String print_triangle(int num) {
        String ret = "";
        for(int i = 0; i < num; i++) {
            for(int j = i+1; j != 0; j--) { ret += "*"; }
            ret += "\n";
        }

        return ret;
    }
}
