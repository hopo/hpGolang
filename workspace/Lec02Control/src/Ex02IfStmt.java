import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Ex02IfStmt {

	public static void main(String[] args) throws IOException {
		// 과제)
		// 국어, 영어, 수학 3과목의 점수를 입력받아 총점, 평균 등급을 구하시오.
		// 95점이상 A+, 90점이상 A, 85점 이상 B+, 80점 이상 B
		// 75점이상 C+, 70점이상 C, 65점 이상 D+, 60점 이상 B, 60점 미만 F
		
		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

		int total = 0;
		System.out.print("*국어 : ");
		total += Integer.parseInt(reader.readLine());
		System.out.print("*영어 : ");
		total += Integer.parseInt(reader.readLine());
		System.out.print("*수학 : ");
		total += Integer.parseInt(reader.readLine());

		float average = total/3f;
		String grade = "";
		
		if(average > 94) {
			grade = "A+";
		} else if(average > 89) {
			grade = "A";
		} else if(average > 84) {
			grade = "B+";
		} else if(average > 79) {
			grade = "B";
		} else if(average > 74) {
			grade = "C+";
		} else if(average > 69) {
			grade = "C";
		} else if(average > 64) {
			grade = "D+";
		} else if(average > 59) {
			grade = "D+";
		} else {
			grade = "F";
		}
		
		System.out.printf("평균: %.2f, %s", average, grade);
	}

}
