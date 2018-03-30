 package arraytest;

public class ArrExtending {
    public static void main(String args[]) {
        Arrlist list = new Arrlist();
        int num = 10;

        for(int i = 0; i < num; i++) {
            list.add(i);
            System.out.print(i + ", ");
        }
    }
}

class Arrlist {
    private int size = 0;
    private Object[] oarr = new Object[10];
    
    public boolean add(Object e) {
        oarr[size] = e;
        size++;
        return true;
    }
}
