
public class Ex03Operator {
	public static void main(String[] args) {
		
		// 식별자(if not Keyword)
		//클래스 : 첫글자는 대문자.
		//메서드, 변수 : 소문자.
		//식별자 이름이 긴 경우. Camel 표기법으로 작성. 예) Integer.parseInt();
		//상수 : final int MAX_NUM = 100; 
		
		// *** Increment operator (++, --)
		int x = 10;

//		int y = x++; // y = x; x += 1;
		int y = ++x; // x += 1; y = x;
		
		System.out.println("x = " + x);
		System.out.println("y = " + y);
		

		
		// *** Arithmetic operator (+, -, *, /, %)
		// 정수형 데이터 연산시 주의사항
		// 1) int형 데이터 보다 작은 정수형 자료형은 int형으로 변환되어 수행.
		// 2) int형 보다 큰 자료형은 큰 자료형으로 형변환 되어서 연산 수행.
		// 3) 실수와 정수와의 연산은 실수로 변환후 연산 수행.

		byte a = 10;
		byte b = 20;
		
		byte c = (byte)(a + b); // (int)a + (int)b
//		int c = a + b;
		System.out.println("c = " + c);
		
		int iNum = 10;
		long lNum = 20;
		
		int result = (int)(iNum + lNum); // (long)iNum + (long)lNum
//		long result = iNum + lNum;
		System.out.println("result = " + result);
		
		float f = 3.0F;
		long result2 = (long)(lNum * f); // (float)lNum * (float)f;
//		float result2 = lNum * f;
		System.out.println("result2 = " + result2);
		
		int n = 234985;
		System.out.println("n % 3 = " + (n % 7)); // 0 ~ 6
		System.out.println("n % 2 = " + (n % 2)); // Odd(1) or Even(0)
		
		
		// *** Shift operator (>>, <<, >>>)
		x = -120;
		System.out.println("x >> 3 = " + (x >> 3)); // 15, 120/8, 120/2^3
		System.out.println("x << 3 = " + (x << 3)); // 960, 120*8, 120*2^3
		System.out.println("x >>> 3 = " + (x >>> 3)); // >>>: unsigned right shift
		
		
		// *** Compare operator (>, <, >=, <=, ==, !=)
		x = 10;
		y = 12;
		
		System.out.println("x > y = " + (x > y));
		System.out.println("x < y = " + (x < y));
		System.out.println("x <= y = " + (x <= y));
	
		/*
		double d = 2;
		System.out.println(Math.pow(d, 0.5));
		System.out.println(Math.sqrt(d));
		*/
		
		
		// *** Logical operator (&&, ||, !)
		// &&(and)
		// T && T = T
		// T && F = F
		// F && T = F
		// F && F = F

		// ||(or)
		// T || T = T
		// T || F = T
		// F || T = T
		// F || F = F
		
		// !(not)
		// !T = F
		// !F = T
		
		x = 10;
		y = 12;
		System.out.println("(x < y) && (x > y) = " + ((x < y) && (x > y)));
		System.out.println("(x < y) && ((x+3) > y) = " + ((x < y) && ((x+3) > y)));
		System.out.println("(x < y) || (x > y) = " + ((x < y) || (x > y)));
		

		// *** 대입 operator
		x = 10;
//		x = x + 20;
		x += 20;
		System.out.println("x = " + x);
		
		
		// *** Bit operator
		x = 170;
		System.out.println("x & 15 = "+ (x & 15)); // 10, 0xaa & 0x0f = 0x0a, mask
		System.out.println("x | 15 = "+ (x | 15)); // 175, 0xaa | 0x0f = 0xaf
		System.out.println("x ^ 15 = "+ (x ^ 15)); // 165, 0xaa ^ 0x0f = 0xa5
		System.out.println("~x = "+ (~x)); // -171,  1의 보수구하기, 비트반전
		System.out.println("~x = "+ ((~x)+1)); // -170, 2의 보수구하기
		
		
		// &&, & 비교
		// 1) F && T -> (F &&) T
		// 2) F & T  -> (F & T)
		// 뒤 항에서 연산을 해야하는 경우는 2번
		
		
		// Unary operator (Conditional operator)
		char gender = 'M'; // M,F
		System.out.println("너어는 "+ (gender == 'M' ? "남자" : "여자") + "!!!");

		
		// *** 과제) 다음의 수가 양수, 음수, 0인지 판별하시오.
		// solution01
		int num = -120;
		String plusminus1 = (num == 0)? "영(0)" : (num > 0)? "양수" : "음수";
		System.out.println(plusminus1);

		/*
		// solution02
		String plusminus2 = (num == 0)? "영(0)" : ((num >>> 31) == 0)? "양수" : "음수";
		System.out.println(plusminus2);

		// solution03
		String plusminus3 = (num == 0)? "영(0)" : ((num & 0x80000000) == 0)? "양수" : "음수";
		System.out.println(plusminus3);
		*/
	
		// Overflow exam
		int oft = Integer.MAX_VALUE;
		System.out.printf("%x", oft);
	}

}
