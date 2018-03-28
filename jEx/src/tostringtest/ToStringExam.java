package tostringtest;

public class ToStringExam {
    public static void main(String args[]) {
        Aobj ao = new Aobj("Show me the ");
        System.out.println(ao);
    }
}

class Aobj {
    private String str = "";
    public Aobj(String str) {
        this.str = str;
    }
    public String toString() {
        return str + "Money!";
    }
}
        
