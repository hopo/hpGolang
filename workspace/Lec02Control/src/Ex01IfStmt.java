
public class Ex01IfStmt {

	public static void main(String[] args) {
		
		char gender = 'M'; // M, F
		
		System.out.print("당신은 ");
		if(gender == 'M') {
			System.out.println("남성입니다.");
		} else if(gender == 'F') {
			System.out.println("여성입니다.");
		} else {
			System.out.println("성별은 T/F중 하나를 입력하세요.");
		}
		
		System.out.println("프로그램 종료.");
		
	}
}
