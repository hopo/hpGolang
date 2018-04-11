import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Exam0411 {

	public static void main(String[] args) throws IOException {

		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		// 과제2> 두 개의 숫자를 입력받아 두 수 사이의 합을 구하시오.
		// 단, 프로그램의 수행 횟수는 항상 일정하게 작성하시오.
		// 예시)
		// 1 ~ 10 합 : 55
		// 1 ~ 100 합 : 5050
		// 1 ~ 10000 합 : 50005000

		int numA, numB, tmp, total = 0;
		while(true) {
			System.out.println(">> 두 수를 입력하시오. (-1:quit) ");		
			System.out.print("시작? ");		
			numA = Integer.parseInt(br.readLine());
			System.out.print("끝? ");		
			numB = Integer.parseInt(br.readLine());
			
			if(numA == -1 || numB == -1) {
				System.out.println("_TERMINATE!");
				break;
			}
			
			if(numA > numB) {
				tmp = numA;
				numA = numB;
				numB = tmp;
			}

			 for(int i = numA; i < numB+1; i++) {
				 total += i;
			 }

			 System.out.printf(" %d ~ %d 합: %d\n", numA, numB, total);
			 System.out.println();
			 numA = numB = -1;
			 total = 0;
		}
	}

}
