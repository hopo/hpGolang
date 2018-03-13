// ing
package level01;

public class NoContinous {
    public static void main(String args[]) {
        NoContinous nc1 = new NoContinous();
        String str1 = "13303";
        String ex1 = nc1.no_continuous(str1); // "1 3 0 3"
    }

    String no_continuous(String str) {
        System.out.println(str); 
        return "-1";
    }
}

/*
def no_continuous(s):

    ret = []
    for i in range(len(s)):
        if s[i-1] != s[i]:
        	ret.append(s[i])
    return ret

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( no_continuous( "133303" ))
*/
