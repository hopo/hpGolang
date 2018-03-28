package level01;

public class WaterMelon {
    public static void main(String args[]) {
        int n1 = 3;
        String ex1 = watermelon(n1); // "수박수"
        System.out.println(ex1);

        int n2 = 4;
        String ex2 = watermelon(n2); // "수박수박"
        System.out.println(ex2);
    }

    static String watermelon(int n) {
        String ret = "";
        for(int i = 1; i < n+1; i++) {
            if(i%2 == 1) { ret += "수"; }
            else { ret += "박"; }
        }

        return ret;
    }
}
