import java.io.IOException;

public class Ex05StdInOut {

	public static void main(String[] args) throws IOException {
		// 표준 입력(키보드로 부터)
		
		// read() return type is int. sign 출력을 위해서 리턴 타입이 int

		System.out.print("*입력 : ");

		// 1.하나의 문자(byte) 입력 받기
//		int x = System.in.read();	// 입력받은 char 정보(0~255) return, 1byte라서 unicode를 읽지는 못한다
//		System.out.println("받은 데이터 : " + (char)x);
//		int x2 = System.in.read() - 48; // -48 or -'0'
//		System.out.println("1.System.in.read() 받은 데이터 : " + x2);


		// 2.여러 문자 입력 받기
//		byte[] data = new byte[20];
//		System.in.read(data);	// parameter byte[]. 입력받은 갯수를 return
//		String str = new String(data); // byte[] -> String
//		System.out.println("2.System.in.read(byte[]) 받은 데이터 : " + str);
		
		
		// 3.하나의 문자(char) 입력받기
//		InputStreamReader isr = new InputStreamReader(System.in);
//		int x3 = isr.read(); // 2byte.
//		System.out.println("3.using InputStreamReader 받은 데이터 : " + (char)x3);

		// 4.여러 문자 입력 받기
//		InputStreamReader isr = new InputStreamReader(System.in);
//		char[] cbuf = new char[20];
//		int x4 = isr.read(cbuf); // 갯수 return
//		System.out.println("4.using InputStreamReader(char[]) 받은 데이터 : " + new String(cbuf));
		
		// 5.한 줄(line)씩 데이터 입력 받기
//		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
//		String str = reader.readLine(); // return String type;
//		System.out.printf("5.받은 String : %s (size:%d)", str, str.length());
	}
	
}