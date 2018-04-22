
def sumDivisor(num): # (int) int
    return sum([0 if num%n else n for n in range(1, num+1)])

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print(sumDivisor(12)) # 28, 1+2+3+4+6+12
    print(sumDivisor(4)) # 7, 1+2+4
