package level01;

public class Fibonacci {
    public static void main(String args[]) {
    	int n1 = 10;
	    int ex1 = fibonacci(n1); // 55
	    System.out.println(ex1);
    }

   static int fibonacci(int n) {
    	int a = 0, b = 1, c = 0;
    	for(int i = 0; i < n-1; i++) {
    		c = a+b;
    		a = b;
    		b = c;
    	}
    
    	return c;
    }
}