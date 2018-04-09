import java.io.IOException;

public class Ex05StdInOut {
	public static void main(String[] args) throws IOException {
		// 표준 입력(키보드로 부터)
		
		// read() return type is int. sign 출력을 위해서 리턴 타입이 int

		System.out.print("입력 : ");

		// 1.하나의 문자 입력 받기
//		int x = System.in.read();	// 입력받은 char 정보(0~255) return
//		System.out.println("받은 데이터 : " + (char)x);
		
//		int x2 = System.in.read() - 48;
//		System.out.println("받은 데이터 : " + x2);

		// 2.여러 문자 입력 받기
		byte[] data = new byte[20];
		System.in.read(data);	// parameter byte[]. 입력받은 char의 수를 return
		String str = new String(data); // byte[] -> String
		System.out.println("받은 데이터 : " + str);
		

	}
	
}