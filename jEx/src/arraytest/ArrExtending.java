package arraytest;

public class ArrExtending {
    public static void main(String args[]) {
        Arrlist list = new Arrlist();
        int num = 12;

        for(int i = 0; i < num; i++) {
            list.add(i);
            System.out.print(i + ", ");
        }
    }
}

class Arrlist {
    private int size = 0;
    private Object[] oarr = new Object[10];
    private Object[] temp;
    
    public boolean add(Object e) {
        if(size == oarr.length) {
            for(int i = 0; i < size; i++) {
                temp = new Object[size+1];
                temp[i] = oarr[i];
            }
            oarr = new Object[size+1];
            for(int i = 0; i < size+1; i++) {
                oarr[i] = temp[i];
            }
        }

        oarr[size] = e;
        size++;
        return true;
    }
}
