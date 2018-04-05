
public class Ex03Operator {
	public static void main(String[] args) {
		
		// *** increment operator (++, --)
		int x = 10;

//		int y = x++; // y = x; x += 1;
		int y = ++x; // x += 1; y = x;
		
		System.out.println("x = " + x);
		System.out.println("y = " + y);
		
		
		// *** arithmetic operator (+, -, *, /, %)
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
		
	}

}
