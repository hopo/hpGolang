import java.util.Arrays;

// package arraytest;

public class Extending {
    public static void main(String args[]) {
        Alist alist = new Alist();
        for(int i = 0; i < 10; i++) {
            alist.add(i);
        }

        // printing
        for(int i = 0; i < alist.length; i++) {
            System.out.print(alist[i] + " ");
        }
    }
}

class Alist {
    private int size = 0;
    private Object[] oarr = new Object[10];
    
    public boolean add(Object e) {
        oarr[size++] = e;
        return true;
    }
}
