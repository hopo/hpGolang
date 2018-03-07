def string_middle(str):
    return len(str)&1 and str[len(str)//2] or str[len(str)//2-1:len(str)//2+1]

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(string_middle("power"))
