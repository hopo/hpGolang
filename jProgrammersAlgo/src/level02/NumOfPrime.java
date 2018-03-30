package level02;

public class NumOfPrime {
	public static void main(String[] args) {
		int num1 = 10;
		int ex1 = numofprime(num1); // 4, [2, 3, 5, 7]
		System.out.println(ex1);

		int num2 = 5;
		int ex2 = numofprime(num2); // 3, [2, 3, 5]
		System.out.println(ex2);
	}
	
	static int numofprime(int num) {
		int p = 0;
		int[] primes = new int[32];
		
		for(int i = 2; i < num+1; i++) {
			// i % each prime != 0. then i is prime
			// pchecker function
			if(pchecker(i, primes)) {
				primes[p++] = i;
			}
		}

		// p is length of array primes
		return p;
	}
	
	static boolean pchecker(int smpl, int[] iarr) {
		boolean ret = true;
		
		for(int i: iarr) {
			// rear of array. not assigned value is 0;
			if(i == 0) { break; }

			// smpl % i == 0. then smpl is not prime number;
			if(smpl%i == 0) {
				ret = false;
				break;
			}
		}

		return ret;
	}
}
