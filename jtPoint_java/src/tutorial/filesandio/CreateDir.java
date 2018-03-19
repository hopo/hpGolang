package tutorial.filesandio;

import java.io.File;

public class CreateDir {
    public static void main(String args[]) {
        String dirname = "/home/hp/workspace/tmp/dirtest";
        File d = new File(dirname);

        // Create directory now.
        d.mkdir();
    }
}
