// ing
package level01;

public class NumPY {
    public static void main(String args[]) {
        String str1 = "pPooyY";
        boolean ex1 = numpy(str1);
        System.out.println(ex1);

        /*
        String cstr2 = "Pyy";
        boolean ex2 = numpy(str2);
        System.out.println(ex2);
        */
    }

    public static boolean numpy(String str) {
        boolean ret = false;
        char[] cst = str.toCharArray();
        System.out.println(cst[0]);

        
        return ret;
    }
}

/*
func main() {
	ex1 := numPY("pPoooyY") // true
	ex2 := numPY("Pyy")     // false
	fmt.Println(ex1)
	fmt.Println(ex2)
}

func numPY(s string) bool {
	var ppp, yyy int
	for _, v := range s {
		switch string(v) {
		case "p":
			ppp++
		case "y":
			yyy++
		}
	}
	return ppp == yyy
}
*/
