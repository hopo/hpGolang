import java.io.IOException;

public class Ex04Switch {
	public static void main(String[] args) throws IOException {
	
		System.out.println("*SELECT (1.Milk, 2.Sugar, 3.Black, 4.Yoolmu): ");
		int x = System.in.read(); // '2'
//		int x = System.in.read() - 48; // '2'-48 => 2
		
		System.out.print("You selected ");
		switch(x) {
		case '1':
			System.out.println("Milk Coffee.");
			break;
		case '2':
			System.out.println("Sugar Coffee.");
			break;
		case '3':
			System.out.println("Black Coffee.");
			break;
		case '4':
			System.out.println("Tea Yoolmu.");
			break;
		default:
			System.out.println("WRONG NUMBER.");
		}
		
	}
}
