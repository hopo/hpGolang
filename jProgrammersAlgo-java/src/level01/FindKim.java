package level01;

public class FindKim {
    public static void main(String args[]) {
        String[] fnames1 = {"Queen", "Tod", "Kim"};
        int ex1 = findkim(fnames1); //2
        System.out.println(ex1);

        String[] fnames2 = {"Gim", "Obama", "xi"};
        int ex2 = findkim(fnames2); //-1
        System.out.println(ex2);
    }

    static int findkim(String[]  fnames) {
        int ret = -1;
        for(int i = 0; i < fnames.length; i++){
            if(fnames[i] == "Kim"){ ret = i; }
        }
        return ret;
    }    
}
