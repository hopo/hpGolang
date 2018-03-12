package level01;

public class SumDivisor {
    public static void main(String args[]) {
        SumDivisor sd1 = new SumDivisor();
	    int ex1 = sd1.sumDivisor(10); // 18, [1, 2, 5, 10] 
        System.out.println(ex1);

        SumDivisor sd2 = new SumDivisor();
	    int ex2 = sd2.sumDivisor(12); // 28, [1, 2, 3, 4, 6, 12] 
        System.out.println(ex2);
    }

    public int sumDivisor(int num) {
	    int ret = 0;	

	    /* Yaksoo of n. total sum */ 
	    for(int i = 1; i < num+1; i++) {
		    if(num%i == 0) { ret += i; }
	    }

	    return ret;
    }
}
