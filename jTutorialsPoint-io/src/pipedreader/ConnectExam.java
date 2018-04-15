package pipedreader;

import java.io.IOException;
import java.io.PipedReader;
import java.io.PipedWriter;

public class ConnectExam {
    public static void main(String args[]) {
        
        // create a new Piped writer and reader
        PipedWriter writer = new PipedWriter();
        PipedReader reader = new PipedReader();

        try {
            
            // connect the reader and the writer
            reader.connect(writer);

            // write somethin
            writer.write(70);
            writer.write(71);

            // read what we wrote
            for(int i = 0; i < 2; i++) {
                System.out.println("" + (char)reader.read());
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

