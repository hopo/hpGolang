
public class Ex04StdInOut {

	public static void main(String[] args) {

		/*
		// *** Standard print (표준 출력)
		System.out.println("표준 출력 예제");
		System.out.print("안녕하세요. ");
		System.out.println("표준 출력 메서드입니다.");

		// printf exam
		int num = 7840;
		System.out.println("10진수 : " + num + " 입니다.");
		System.out.printf("10진수 : %5d 입니다.\n", num);
		System.out.printf("10진수 : %-5d 입니다.\n", num);
		System.out.printf("10진수 : %05d 입니다.\n", num);

		System.out.printf("8진수 : 0o%o 입니다.\n", num);
		System.out.printf("16진수 : 0x%x 입니다.\n", num);
		System.out.printf("16진수 : 0x%X 입니다.\n", num);
		
		double num2 = 123.456789;
		System.out.printf("실수형 : %f 입니다.\n", num2);
		System.out.printf("실수형 : %6.2f 입니다.\n", num2); // 소숫점 아래 2자리, 전체 6자리 
		

		String str = "String String String";
		System.out.printf("String : %s 입니다.\n", str);

		char ch = 'r';
		System.out.printf("char : %s 입니다.\n", ch);

		boolean bl = 10 < 20;
		System.out.printf("boolean : %b 입니다.\n", bl);
		*/
		
		
		// Example: 이름 나이 성별 체중 신장 결혼 전화 주소
		String name = "남숙희";
		int age = 34;
		char gender = 'W';
		float weight = 53.8F;
		double height = 160.2;
		boolean married = true;
		String phone = "011-8321-3698";
		String address = "오사카부 오사카시 이토노구";

		System.out.print("=============== Print Private Infomation ===============\n");
		System.out.print("이름\t 나이\t 성별\t 체중\t 신장\t 결혼\t 전화\t\t 주소\n");
		System.out.printf(
			"%s\t %d\t %c\t %5.2f\t %6.2f\t %b\t %s\t %s\n",
			name, age, gender, weight, height, married, phone, address
		);
		
	}

}
