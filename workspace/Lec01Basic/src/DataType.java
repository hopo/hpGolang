class DataType {

	public static void main(String[] args) {

		// 논리형, boolean
		boolean bool;	// variable declare
		bool = true;	// assign
		System.out.println("bool = " + bool); // using variable

		bool = 10 < 20;	// true
		System.out.println("bool = " + bool);

		// byte, int, short, long
		byte b = 127;
		System.out.println("\nbyte = " + b); 

		int n = 2140000000;
		System.out.println("int = " + n); 

		long m = 3140000000L;
		System.out.println("long = " + m); 

		// character
		char ch = 'A';	// 65
		System.out.println("\nch = " + ch); 

		ch = 66;	// 'B'
		System.out.println("ch = " + ch); 

		ch = '가';
		System.out.println("ch = " + ch); 

		ch = '\uAC01';	// using unicode
		System.out.println("ch = " + ch); 

		// float, double
		float f = 3.14F;
		System.out.println("\nfloat = " + f);

		double d = 12.254;
		System.out.println("double = " + d);

		// String
		String name = "HONG GIL DONG";
		System.out.println("\nString = " + name);
	}

}
