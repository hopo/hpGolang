package bufferdreader;

import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.InputStreamReader;

public class readExam02 {
	
	public static void main(String[] args) throws Exception {
		
		InputStream is = null;
		InputStreamReader isr = null;
		BufferedReader br = null;
		
		try {
			
			// open input stream test2.txt for reading purpose.
			is = new FileInputStream("/home/pc33/hpWorks/javatest/test2");
			
			// create new input stream reader
			isr = new InputStreamReader(is);
			
			// create new buffered reader
			br = new BufferedReader(isr);
			
			// creates buffer
			char[] cbuf = new char[is.available()];
			
			// reads characters to buffer, offset 2, len 10
			br.read(cbuf, 2, 10);
			
			// for each character in the buffer
			for(char c: cbuf) {
				
				// if char is empty
				if(c == (char)0) {
					c = '*';
				}
				
				// prints characters
				System.out.print(c);
			}
			
		} catch(Exception e) {
			e.printStackTrace();
			
		} finally {
			
			// release resources associated with the streams
			if(is != null) { is.close(); }
			if(isr != null) { is.close(); }
			if(br != null) { is.close(); }
		}
	}

}
