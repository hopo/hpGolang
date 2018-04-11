import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class DecimalConvExam {

	public static void main(String[] args) throws IOException {

		// 과제1> 10진수를 입력 받아 2진수, 8진수, 16진수로 출력하시오.
		// 예시)
		// 10진수 : 170
		// === 출력결과 ===
		// 2진수 : 00000000000000000000000010101010
		// 8진수 : 00000000252
		// 16진수 : 000000AA

		// 입력받기 위해 BufferdReader 참조형 변수 br 생성.
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		// variable declaration
		int num, decimal;

		// 최대 10회 질문.
		for(int i = 0; i < 10; i++) {

			// input number
			System.out.print(">>> INPUT NUMBER (-1:quit) ? ");

			// input number is decimal. 
			decimal = Integer.parseInt(br.readLine());
			
			// if input number is -1 then program terminate.
			if (decimal == -1) {
				System.out.println("Program Terminated!");
				break;
			}

			// 2진수 변환.
			num = decimal;	// num using operand
			String binary = "";
			for(int j = 0; j < 32; j++) {
				if(num == 0) {			// num이 0이 되었다면 나머지 자리에는 0으로 채운다.
					binary = "0" + binary;
					continue;
				}
				binary = (num%2) + binary;

				num /= 2;
			}

			// 8진수 변환.
			num = decimal;	// num using operand
			String octal = "";
			for(int j = 0; j < 11; j++) {
				if(num == 0) {			// num이 0이 되었다면 나머지 자리에는 0으로 채운다.
					octal = "0" + octal;
					continue;
				}
				octal = (num%8) + octal;

				num /= 8;
			}

			// 16진수 변환.
			num = decimal;	// num using operand
			String hexadecimal = "";
			for(int j = 0; j < 8; j++) {
				if(num == 0) {		 // num이 0이 되었다면 나머지 자리에는 0으로 채운다.
					hexadecimal = "0" + hexadecimal;
					continue;
				}

				// num/16의 나머지가 10~15 라면 A~F.
				if (num%16 > 9) {
					hexadecimal = (char)(num%16 + 55) + hexadecimal;
				} else {
					hexadecimal = (num%16) + hexadecimal;
				}

				num /= 16;
			}
			
			// print results
			System.out.printf("10진수: %d\n", decimal);
		    System.out.print("=== 출력결과 ===\n");
			System.out.printf(" 2진수: %s\n", binary);
			System.out.printf(" 8진수: %s\n", octal);
			System.out.printf("16진수: %s\n", hexadecimal);
		    System.out.println();
		}
	}

}
