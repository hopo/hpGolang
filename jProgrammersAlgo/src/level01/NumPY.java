package level01;

public class NumPY {
    public static void main(String args[]) {
        String str1 = "pPooyY";
        boolean ex1 = numpy(str1); // true
        System.out.println(ex1);

        String str2 = "Pyy";
        boolean ex2 = numpy(str2); // false
        System.out.println(ex2);
    }

    static boolean numpy(String str) {
        int ppp = 0, yyy = 0;
        for(int i = 0; i < str.length(); i++) {
            switch (str.charAt(i)) {
                case 'P':
                case 'p':
                    ppp++;
                    break;
                case 'Y':
                case 'y':
                    yyy++;
                    break;
            }
        }

        return ppp == yyy;
    }
}
