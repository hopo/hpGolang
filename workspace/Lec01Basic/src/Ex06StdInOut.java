import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Ex06StdInOut {
	public static void main(String[] args) throws IOException {

		// "이름\t 나이\t 성별\t 체중\t 신장\t 결혼\t 전화\t\t 주소\n"		
		System.out.println("[[ 입력하시오 ]]");
		System.out.print("*이름: "); 
		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
		String name = reader.readLine();

		System.out.print("*나이(int): "); 
		BufferedReader reader2 = new BufferedReader(new InputStreamReader(System.in));
		int age = Integer.parseInt(reader2.readLine());

		System.out.print("*성별(m/f): "); 
		BufferedReader reader3 = new BufferedReader(new InputStreamReader(System.in));
		String gender = reader3.readLine();

		System.out.print("*체중(double): "); 
		BufferedReader reader4 = new BufferedReader(new InputStreamReader(System.in));
		double weight = Double.parseDouble(reader4.readLine());

		System.out.print("*신장(double): "); 
		BufferedReader reader5 = new BufferedReader(new InputStreamReader(System.in));
		double height = Double.parseDouble(reader5.readLine());

		System.out.print("*결혼(true/false): "); 
		BufferedReader reader6 = new BufferedReader(new InputStreamReader(System.in));
		boolean married = Boolean.parseBoolean(reader6.readLine());

		System.out.print("*전화(String): "); 
		BufferedReader reader7 = new BufferedReader(new InputStreamReader(System.in));
		String phone = reader7.readLine();

		System.out.print("*주소(String): "); 
		BufferedReader reader8 = new BufferedReader(new InputStreamReader(System.in));
		String address = reader8.readLine();
		
		
		System.out.print("=============== Print Private Infomation ===============\n");
		System.out.print("이름\t 나이\t 성별\t 체중\t 신장\t 결혼\t 전화\t\t 주소\n");
		System.out.printf(
			"%s\t %d\t %s\t %5.2f\t %6.2f\t %b\t %s\t %s\n",
			name, age, gender, weight, height, married, phone, address
		);	
		
	}
}
