public class PYExam {
	public static void main(String[] args) {
		int yr1 = 2000; // Y
		int yr2 = 2016; // Y
		int yr3 = 2100; // P
		int yr4 = 2018; // P
		
		String ex1 = gene(yr1);
		String ex2 = gene(yr2);
		String ex3 = gene(yr3);
		String ex4 = gene(yr4);
		
		System.out.printf("%s ", ex1);
		System.out.printf("%s ", ex2);
		System.out.printf("%s ", ex3);
		System.out.printf("%s ", ex4);
	}
	
	static String gene(int year) {
		
		boolean bl4 = year%4 == 0;
		boolean bl100 = year%100 == 0;
		boolean bl400 = year%400 == 0;
		
		/*
		// s01)
		String ret = "";
		if(bl4) {
			if(bl100) {
				if(bl400) {
					ret = "Y";
				} else {
					ret = "P";
				}
			} else {
				ret = "Y";
			}
		} else {
			ret = "P";
		}
		*/

		// s02)
		String ret = "P";
		if(bl4 && !bl100 || bl400) { ret = "Y"; }
		
		return ret;
	}
}
