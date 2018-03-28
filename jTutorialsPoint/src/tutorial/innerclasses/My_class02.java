package tutorial.innerclasses;

class Outer_Demo02{
    // private variable of the outer class
    private int num = 175;

    // inner class
    public class Inner_Demo02 {
        public int getNum() {
            System.out.print("This is the getnum method of the inner class ");
            return num;
        }
    }
}

public class My_class02 {

    public static void main(String args[]) {
        // Instantiating the outer class
        Outer_Demo02 outer = new Outer_Demo02();

        // Instantiating the inner class
        Outer_Demo02.Inner_Demo02 inner = outer.new Inner_Demo02();
        System.out.println(inner.getNum());
    }
}
