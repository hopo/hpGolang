public class Ex02Person {

	public static void main(String[] args) {

		// 과제: 이름 나이 성별 체중 신장 결혼 전화 주소
		String name = "riceboy";
		int age = 34;
		char gender = 'M';
		float weight = 63.8F;
		double height = 180.2;
		boolean married = true;
		String phone = "011-8321-3698";
		String address = "Oryu-dong Jung-gu Daejeon-city";

		System.out.println(" ========== Print Info ========== ");
		System.out.println("이름: " + name);
		System.out.println("나이: " + age);
		System.out.println("성별: " + gender);
		System.out.println("체중: " + weight);
		System.out.println("신장: " + height);
		System.out.println("결혼: " + married);
		System.out.println("전화: " + phone);
		System.out.println("주소: " + address);
	}

}