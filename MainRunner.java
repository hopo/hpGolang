// 실행을 위한 코드입니다. 수정하지 마세요.
public class MainRunner 
{
    public static void main(String[] args)
    {
        ArrayList arraylist = new ArrayList();
        for(int i=0; i<100; i++){
            arraylist.addLast(i);
        }
        return ;
    }
}

class ArrayList {
    private int size = 0;
    private Object[] elementData = new Object[100];
 
    public ArrayList() {}
     
    public boolean addLast(Object element) {
        if(size == elementData.length) {
            elementData = new Object[size*10];
        }
        elementData[size] = element;
        size++;
        return true;
    }
}

