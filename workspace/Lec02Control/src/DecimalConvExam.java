import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class DecimalConvExam {

	public static void main(String[] args) throws IOException {

		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		// 과제1> 10진수를 입력 받아 2진수, 8진수, 16진수로 출력하시오.
		// 예시)
		// 10진수 : 170
		// === 출력결과 ===
		// 2진수 : 00000000000000000000000010101010
		// 8진수 : 00000000252
		// 16진수 : 000000AA

		int num, tmp;
		while (true) {
			System.out.print(">>> INPUT NUMBER (-1:quit) ? ");
			num = Integer.parseInt(br.readLine());
			tmp = num;
			if (num == 0) {
				System.out.println("_TERMINATE!");
				break;
			}

			num = tmp;
			String binary = "";
			for(int i = 0; i < 32; i++) {
				if(num == 0) {
					binary = "0" + binary;
					continue;
				}
				binary = (num%2) + "" + binary;
				num /= 2;
			}

			num = tmp;
			String octal = "";
			for(int i = 0; i < 11; i++) {
				if(num == 0) {
					octal = "0" + octal;
					continue;
				}
				octal += (num%8) + "";
				num /= 8;
			}

			num = tmp;
			String hexadeci = "";
			for(int i = 0; i < 8; i++) {
				if(num == 0) {
					hexadeci = "0" + hexadeci;
					continue;
				}
				if (num%16 > 9) {
					hexadeci += (char)(num % 16 + 55) + "";
				} else {
					hexadeci += (num%16) + "";
				}
				num /= 16;
			}
			
			System.out.printf("10진수: %d\n", tmp);
		    System.out.println("=== 출력결과 ===");
			System.out.printf(" 2진수: %s\n", binary);
			System.out.printf(" 8진수: %s\n", octal);
			System.out.printf("16진수: %s\n", hexadeci);
		    System.out.println();
			num = 0;
		}
	}

}
