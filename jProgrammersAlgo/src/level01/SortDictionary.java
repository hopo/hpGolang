package level01;

import java.util.Map;
import java.util.Arrays;
import java.util.HashMap;

public class SortDictionary {
	public static void main(String args[]) {
        HashMap<String, Integer> ncard1 = new HashMap<>();
        ncard1.put("정진원", 88);
        ncard1.put("이하나", 97);
        ncard1.put("김철수", 78);

        String ex1 = sortdictionary(ncard1); // ["김철수" = 78], ["이하나" = 97], ["정진원" = 88]
        System.out.println(ex1);
	}

    static String sortdictionary(Map<String, Integer> ncard) {
    	int s = 0;
        String k = "", dict = "", ret = "";
    	String[] keyarray = new String[ncard.size()];

        for(Map.Entry<String, Integer> m : ncard.entrySet()) {
        	keyarray[s++] = m.getKey();
        }
        
        Arrays.sort(keyarray);

        for(int i = 0; i < keyarray.length; i++) {
        	k = keyarray[i];
        	dict = String.format("[\"%s\" = %d]", k, ncard.get(k));
        	ret += dict;
        	if(i != keyarray.length-1) {
        		ret += ", ";
        	}
        }
        
        return ret;
    }
}
