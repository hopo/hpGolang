package level01;

public class StrToInt {
    public static void main(String args[]) {
        String str1 = "-1234";
        int ex1 = strtoint(str1); // -1234
        System.out.println(ex1);

        String str2 = "a1234";
        int ex2 = strtoint(str2); // error
        System.out.println(ex2);
    }

    static int strtoint(String str) {
        int ret = 0;
        try {
            ret = Integer.parseInt(str);
        } catch(Exception e) {
            e.printStackTrace();
        } 

        return ret;
    }
}
