package level01;

// solution02: using StringBuffer

public class HideNumbers02 {
    public static void main(String args[]) {
        String pnum1 = "01033334444";
        String ex1 = hidenumber(pnum1); // "*******4444"
        System.out.println(ex1);

        String pnum2 = "027778888";
        String ex2 = hidenumber(pnum2); // "*****8888"
        System.out.println(ex2);
    }

    static String hidenumber(String pnum) {
    	int slnth = pnum.length()-4;
    	StringBuffer sb = new StringBuffer(pnum);
    	
    	String stars = "";
    	for(int i = 0; i < slnth; i++) { stars += "*"; }
    	
    	return sb.replace(0, slnth, stars).toString();
    }
}
