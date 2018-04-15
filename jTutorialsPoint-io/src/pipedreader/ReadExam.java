package pipedreader;

import java.io.IOException;
import java.io.PipedReader;
import java.io.PipedWriter;

// reader.read(array, bidx, eidx)
public class ReadExam {
    public static void main(String args[]) {
        
        // create a new Piped writer and reader
        PipedWriter writer = new PipedWriter();
        PipedReader reader = new PipedReader();

        try {
            
            // connect the reader and the writer
            reader.connect(writer);

            // write something
            writer.write(70);
            writer.write(71);

            // read into a char array
            char[] b = new char[2];
            reader.read(b, 0, 2);

            // print the char array
            for(int i = 0; i < 2; i++) {
                System.out.println("" + b[i]);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

