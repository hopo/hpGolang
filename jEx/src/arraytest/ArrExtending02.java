// package arraytest;

public class ArrExtending02 {
    public static void main(String args[]) {
        Arrlist list = new Arrlist();
        int num = 1000;

        // for test
        for(int i = 0; i < num; i++) {
            list.addLast(i);
            System.out.print(i + ", ");
        }
    }
}

class Arrlist {
    private int size = 0;
    private int capa = 100;
    private Object[] elementData = new Object[capa];

    public void resize() {
        Object[] newArray = new Object[capa+1];
        for (int i = 0; i < capa; i++) {
            newArray[i] = elementData[i];
        }
        elementData = newArray;
        capa++;
    }

    public boolean addLast(Object element) {
        if (size >= capa) { resize(); }
        elementData[size++] = element;
        return true;
    }
}
