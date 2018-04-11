import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Ex02IfStmt02 {
	public static void main(String[] args) throws IOException {
		// 국어, 영어, 수학 3과목의 점수를 입력받아 총점, 평균 등급을 구하시오.
		// 95점이상 A+, 90점이상 A, 85점 이상 B+, 80점 이상 B
		// 75점이상 C+, 70점이상 C, 65점 이상 D+, 60점 이상 B, 60점 미만 F
		
		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
		
		System.out.print("국어 : ");
		int kor = Integer.parseInt(reader.readLine());
		System.out.print("영어 : ");
		int eng = Integer.parseInt(reader.readLine());
		System.out.print("수학 : ");
		int mat = Integer.parseInt(reader.readLine());
		
		int total = kor + eng + mat;
		float avg = (float)total / 3;
		System.out.println(avg);
		
		String grade ="";
		if(avg >= 90) {
			grade = "A";
		} else if(avg >= 80) {
			grade = "B";
		} else if(avg >= 70) {
			grade = "C";
		} else if(avg >= 60) {
			grade = "D";
		} else {
			grade = "F";
		}
		
		// 1의 자리를 구한다 (avg% 10 >= 5)
		if(avg >= 65 && (avg%10 >= 5) || avg == 100) {
			grade += "+";
		}
	
		System.out.println(">>>>>>>>>> 중간고사 성적표 <<<<<<<<<<");
		System.out.println("국어\t 영어\t 수학\t 총점\t 평균\t 등급");
		System.out.printf("%d\t %d\t %d\t %d\t %.2f\t %s", kor, eng, mat, total, avg, grade);
	}

}
