import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Ex06StdInOut03 {
	public static void main(String[] args) throws IOException {

		// 과제) "이름\t 나이\t 성별\t 체중\t 신장\t 결혼\t 전화\t\t 주소\n"		

		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

		System.out.println("[[ *입력하시오* ]]");

		System.out.print("*이름: "); 
		String name = reader.readLine(); // readLine()은 Enter(Line feed, \n) 기준으로 라인 읽기

		System.out.print("*나이: "); 
		int age = Integer.parseInt(reader.readLine());

		System.out.print("*성별(m/f): "); 
		char gender = (char)(reader.read()); // (char)"m\n", '\n'이 남게된다.
		reader.readLine(); // 그래서 readLine()을 이용하여 '\n' 읽고 날려버리기. linux(\n) window(\n\r)

		System.out.print("*체중: "); 
		float weight = Float.parseFloat(reader.readLine());

		System.out.print("*신장: "); 
		double height = Double.parseDouble(reader.readLine());

		System.out.print("*결혼(true/false): "); 
		boolean married = Boolean.parseBoolean(reader.readLine());

		System.out.print("*전화: "); 
		String phone = reader.readLine();

		System.out.print("*주소: "); 
		String address = reader.readLine();
		
		// data print
		System.out.print("==================== Print Private Infomation ====================\n");
		System.out.print("이름\t 나이\t 성별\t 체중\t 신장\t 결혼\t 전화\t\t 주소\n");
		System.out.printf(
			"%s\t %d\t %c\t %5.2f\t %6.2f\t %b\t %s\t %s\n",
			name, age, gender, weight, height, married, phone, address
		);	
		System.out.print("==================================================================\n");
		
	}
}
