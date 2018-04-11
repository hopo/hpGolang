import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class SumTwoNumExam {

	public static void main(String[] args) throws IOException {

		// 과제2> 두 개의 숫자를 입력받아 두 수 사이의 합을 구하시오.
		// 단, 프로그램의 수행 횟수는 항상 일정하게 작성하시오.
		// 예시)
		// 1 ~ 10 합 : 55
		// 1 ~ 100 합 : 5050
		// 1 ~ 10000 합 : 50005000

		// 입력받기 위해 BufferdReader 참조형 변수 br 생성.
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		int numA, numB, tmp, total = 0;

		// 최대 10회 질문.
		for(int i = 0; i < 10; i++) {
			System.out.println(">> 두 수를 입력하시오. (-1:quit) ");		
			System.out.print(" 시작? ");		
			numA = Integer.parseInt(br.readLine());
			System.out.print("  끝? ")	;
			numB = Integer.parseInt(br.readLine());
			
			// 만약 입력받은 두 수 중 하나가 -1 이면 프로그램 종료.
			if(numA == -1 || numB == -1) {
				System.out.println("Program Terminated!");
				break;
			}
			
			// 만약 처음 입력 받은 수가 뒤에 입력 받은 수 보다 크다면 스왑 해준다.
			if(numA > numB) {
				tmp = numA;
				numA = numB;
				numB = tmp;
			}

			// sum from numA to numB
			 for(int j = numA; j < numB+1; j++) {
				 total += j;
			 }

			 // print result
			 System.out.printf("%d ~ %d 합: %d\n", numA, numB, total);
			 System.out.println();
			 numA = numB = -1;
			 total = 0;
		}
	}

}
