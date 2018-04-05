
/**
 * 
 *  Java Document Comments
 *  
 *  @author pc33
 *  @since 2018
 *  @version 1.8
 *
 */

public class Ex01Casting {
	public static void main(String[] args) {
		
		// 한줄 주석

		/*
		여러줄 주석
		*/
		
		byte byteNum = 120;
		
		int intNum = byteNum;	// Promotion (자동형변환)
		System.out.println("intNum = " + intNum);

		short shortNum = 150;
		byteNum = (byte)shortNum; // Casting - Demotion (강제형변환)
		System.out.println("byteNum = " + byteNum);
		
		char ch = 65;	// 'A'
		shortNum = (short)ch;
		System.out.println("shortNum = " + shortNum);
		
		float floatNum = 3.14F;	// 4byte
		long longNum = (long)floatNum; // 8byte, Casting
		System.out.println("longNum = " + longNum);
		
		String str = "java \"Hello\" \'GoGoGo\'";	// escape
		String path = "c:\\home\\pc14"; // path separate in Windows OS
		System.out.println(str);
		System.out.println(path);
		
		String message = "안녕하세요.\n저의\t이름은\t오대수입니다.";
		System.out.println(message);
		
		String sNum = "120" + 20;	
		System.out.println("sNum = " + sNum);

		int num = Integer.parseInt(sNum) + 20;
		System.out.println("num = " + num);
		
		String strBool = "true"; // string "true"
		boolean bool = Boolean.parseBoolean(strBool);
		System.out.println(bool);
	}
}