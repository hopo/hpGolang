package level01;

import java.util.*;

public class SortDictionary {
	public static void main(String args[]) {
        Map<String, Integer> ncard1 = new HashMap<String, Integer>();
        ncard1.put("김철수", 78);
        ncard1.put("이하나", 97);
        ncard1.put("정진원", 88);

       sortdictionary(ncard1); // ["김철수" = 78], ["이하나" = 97], ["정진원" = 88]
	}

    // sorting?
    public static void sortdictionary(Map<String, Integer> ncard) {
        for(Map.Entry m : ncard.entrySet()) {
            // System.out.print(m);
            // System.out.print(" ");

            System.out.print(m.getKey()+"="+m.getValue());
            System.out.print(" ");
        }
    }
}
