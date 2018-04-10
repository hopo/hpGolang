package bufferdreader;

import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.InputStreamReader;

public class readExam {

	public static void main(String[] args) throws Exception {
		InputStream is = null;
		InputStreamReader isr = null;
		BufferedReader br = null;
		
		try {

			// open input stream test.txt for reading purpose.
			is =  new FileInputStream("/home/pc33/hpWorks/javatest/test.txt");
			
			// create new input stream reader
			isr = new InputStreamReader(is);
			
			// create new buffered reader
			br = new BufferedReader(isr);
			
			int value = 0;
			
			// read to the end of the stream
			while((value = br.read()) != -1) {
				
				// converts int to character
				char c = (char)value;
				
				// prints character
				System.out.printf("%c ", c);
				
			}
			
		} catch(Exception e) {
			e.printStackTrace();
		} finally {
			
			// release resource associated with the streams
			if(is != null) { is.close(); }
			if(isr != null) { isr.close(); }
			if(br != null) { br.close(); }
		}
		
	}
}
