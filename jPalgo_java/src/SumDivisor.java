package level01;

public class SumDivisor {
    public static void main(String args[]) {
	    int ex1 = sumdivisor(10); // 18, [1, 2, 5, 10] 
        System.out.println(ex1);

	    int ex2 = sumdivisor(12); // 28, [1, 2, 3, 4, 6, 12] 
        System.out.println(ex2);
    }

    public static int sumdivisor(int num) {
	    int ret = 0;	

	    /* Yaksoo of n. total sum */ 
	    for(int i = 1; i < num+1; i++) {
		    if(num%i == 0) { ret += i; }
	    }

	    return ret;
    }
}
