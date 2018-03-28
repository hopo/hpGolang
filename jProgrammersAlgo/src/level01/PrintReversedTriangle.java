package level01;

public class PrintReversedTriangle {
    public static void main(String args[]) {
        int n1 = 3;
        String ex1 = printreversedtriangle(n1); // "***\n**\n*\n"
        System.out.println(ex1);

        int n2 = 50;
        String ex2 = printreversedtriangle(n2); // "*****\n****\n***\n**\n*\n"
        System.out.println(ex2);
    }
    
    static String printreversedtriangle(int n) {
        String ret = "";
        StringBuffer sbff = new StringBuffer();
        for(int i = n; i != 0; i--) {
            for(int j = 0; j < i; j++) {
                sbff.append("*");
            }
            sbff.append("\n");
        }

        ret = sbff.toString();
        return ret;
    }
}
