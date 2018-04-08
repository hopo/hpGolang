package level02;

public class JadenCase {
    public static void main(String args[]) {
        String str1 = "3people unFollowed me for the last week"; 
        String ex1 = jdcase(str1); // "3people Unfollowed Me For The Last Week"
        System.out.println(ex1);

        String str2 = "viDeo kIl1ed radio 5taR"; 
        String ex2 = jdcase(str2); // "Video Kil1ed Radio 5tar" 
        System.out.println(ex2);
    }

    static String jdcase(String str) {
        char[] chArr = str.toCharArray();
        boolean flag = true;

        for(int i = 0; i < chArr.length; i++) {
            if(chArr[i] == ' ') {
                flag = true;
                continue;
            }
            if(flag) {
                if(chArr[i] >= 'a' && 'z' >= chArr[i]) {
                    chArr[i] -= 32;
                }
                flag = false;
            } else {
                if(chArr[i] >= 'A' && 'Z' >= chArr[i]) {
                    chArr[i] += 32;
                }
            }
        }
        
        return new String(chArr);
    }
}
