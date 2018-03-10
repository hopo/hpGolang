package ifelse;

public class MainIfelse {
  public static void main(String args[]) {
    int aaa = 0xff;
    int bbb = 0x11;
    if(aaa > bbb) {
      System.out.print("1.aaa bigger than bbb: ");
      System.out.println(aaa);
    }
    else {
      System.out.print("2.bbb bigger than aaa: ");
      System.out.println(bbb);
    }
  }
}
