package strings;

public class StringBuffer {
    public static void main(String args[]) {
        countTo_N_Improved();
    }

    final private static int N = 100;
    final private static int MAX_LENGTH = 30;
    private static String buffer = "";

    private static void countTo_N_Improved() {
        for(int count = 2; count <= N; count += 2) {
            emit(" " + count);
        }
    }

    private static void emit(String nextChunk) {
        if(buffer.length()+nextChunk.length() > MAX_LENGTH) {
            System.out.println(buffer);
            buffer = "";
        }
        buffer += nextChunk;
    }
}
