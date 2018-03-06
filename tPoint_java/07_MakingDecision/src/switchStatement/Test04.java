package switchStatement;

public class Test04 {
	
	public static void main(String args[]) {
		// char grade = args[0].charAt(0);
		char grade = 'C';
		
		switch(grade) {
		case 'A':
			System.out.println("Exceleent!");
			break;
		case 'B':
		case 'C':
			System.out.println("Well done!");
			break;
		case 'D':
			System.out.println("You passed");
		case 'F':
			System.out.println("Better try again");
			break;
		default:
			System.out.println("Invalid grade");
		}
		System.out.println("Yout grade is " + grade);
	}
}
