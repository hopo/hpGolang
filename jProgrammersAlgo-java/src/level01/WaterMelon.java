package level01;

public class WaterMelon {
    public static void main(String args[]) {
        int n1 = 3;
        String ex1 = watermelon(n1); // "수박수"
        System.out.println(ex1);

        int n2 = 4;
        String ex2 = watermelon(n2); // "수박수박"
        System.out.println(ex2);
    }

    static String watermelon(int n) {
        String ret = "";
        for(int i = 1; i < n+1; i++) {
            if(i%2 == 1) { ret += "수"; }
            else { ret += "박"; }
        }

        return ret;
    }
}

/*
func main() {
	ex1 := water_melon(3)
	ex2 := water_melon(4)
	fmt.Println(ex1) // "WMW"
	fmt.Println(ex2) // "WMWM"
}

func water_melon(n int) (ret string) {
	for i := 1; i < n+1; i++ {
		if i%2 == 1 {
			ret += "W"
		} else {
			ret += "M"
		}
	}
	return
}
*/

/*
# https://programmers.co.kr/learn/challenge_codes/108

def water_melon(n):
    # 함수를 완성하세요.
    return ""


# 실행을 위한 테스트코드입니다.
print("n이 3인 경우: " + water_melon(3));
print("n이 4인 경우: " + water_melon(4));
*/
