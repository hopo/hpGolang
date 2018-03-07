def no_continuous(s):
    ret = []
    for i in range(len(s)):
        if s[i-1] != s[i]:
        	ret.append(s[i])
    return ret

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( no_continuous( "133303" ))
