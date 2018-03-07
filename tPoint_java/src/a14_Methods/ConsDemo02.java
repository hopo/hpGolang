package a14_Methods;


public class ConsDemo02 {
	
	public static void main(String args[]) {
		MyClass02 tt1 = new MyClass02(10);
		MyClass02 tt2 = new MyClass02(20);
		System.out.println(tt1.x + " " + tt2.x);
	}

}

// A simple constructor.
class MyClass02 {
	int x;
	
	// Folowing is the constructor
	MyClass02(int i) {
		x = i;
	}
}