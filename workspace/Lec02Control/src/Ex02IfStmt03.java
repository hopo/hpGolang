import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Ex02IfStmt03 {
	public static void main(String[] args) throws IOException {
		
		// 태양력(1년 : 365.2425)
		// 특정 년을 입력 받아 해당 년도가 윤년인지 평년인지 판별하시오.
		// 1. 해당 년이 4로 나누어 떨어지면 (4년마다 발생)
		// 2. 1의 조건을 만족하는 것 중 100으로 나어 떨어지면 평년 (100년 마다 평년.)
		// 3. 2의 조건을 만족하는 것 중, 400으로 나눠 떨어지면 윤년 (400년 마다 윤년)
		
		BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
		
		System.out.println("년도를 입력하세요:");
		int year = Integer.parseInt(reader.readLine());
		
		boolean bl1 = year%4 == 0;
		boolean bl2 = year%100 == 0;
		boolean bl3 = year%400 == 0;
		
		/*
		// solution 1
		String ret = "평";
		if(bl1) {
			if(bl2) {
				if(bl3) {
					ret = "윤";
				}
			} else {
				ret = "윤";
			}
		}
		*/

		// solution 2
		String ret = "평";
		if(bl1 && !bl2 || bl3) {
				ret = "윤";
		}
		
		/*
		// solution 3
		String ret = (bl1 && !bl2 || bl3)? "윤" : "평";
		*/
		
		int leapYearCnt = (year/4) - (year/100) + (year/400);
		System.out.printf("%d년은 %s년 입니다.(윤년의 횟수: %d)", year, ret, leapYearCnt);
		
	}
}