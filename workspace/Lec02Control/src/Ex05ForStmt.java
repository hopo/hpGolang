
public class Ex05ForStmt {

	public static void main(String[] args) {

		// // *-- for statement
		// for (int i = 1; i <= 5; i++) {
		// System.out.printf("Hello, Java! - %d\n", i);
		// }

		// // *-- 1 ~ 10 까지의 합
		// int sum = 0;
		// for (int i = 1; i <= 10; i++) {
		// sum += i;
		// }
		// System.out.println("1 ~ 10 : " + sum);

		// // *-- 1 ~ 100 까지의 3의 배수의 합
		// sum = 0;
		// for (int i = 1; i <= 100; i++) {
		// if (i % 3 != 0) {
		// continue;
		// }
		// sum += i;
		// }
		// System.out.println("1 ~ 100 : " + sum);

		// // *-- A ~ Z 까지 출력
		// for (char ch = 'A'; ch <= 'Z'; ch++) {
		// System.out.println(ch + " ");
		// }

		// // *-- 구구단 출력
		// for (int i = 2; i <= 9; i++) {
		// System.out.printf("*%d단\n", i);
		// for (int j = 2; j <= 9; j++) {
		// System.out.printf("%d x %d = %d\n", i, j, i * j);
		// }
		// }

		// // *-- print **********
		// for (int i = 0; i < 10; i++) {
		// System.out.print("*");
		// }

		// *-- print 5line **********
		for (int j = 0; j < 5; j++) {
			for (int i = 0; i < 10; i++) {
				System.out.print("*");
			}
			System.out.println();
		}

	}
}
