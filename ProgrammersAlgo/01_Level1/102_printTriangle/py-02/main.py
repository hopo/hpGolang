
def printTriangle(num): # (int) str
    s = ""
    for v in range(1, num+1):
        s += ('*' * v) + '\n'
    return s

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print(printTriangle(3)) # "*\n**\n***\n"
