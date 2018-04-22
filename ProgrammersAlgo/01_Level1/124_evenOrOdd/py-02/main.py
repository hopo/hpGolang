
def evenOrOdd(num): # (int) str
    return num&1 and "Odd" or "Even"

#아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print("결과 : " + evenOrOdd(3)) # "Odd"
    print("결과 : " + evenOrOdd(2)) # "Even"
