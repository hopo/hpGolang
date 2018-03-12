def printTriangle(num):
    s = ""
    for v in range(1, num+1):
        s += ('*' * v) + '\n'
    return s

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print( printTriangle(3) )
