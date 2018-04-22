
def sum_digit(number): # (int) str
    return sum(int(v) for v in str(number))

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print("결과 : {}".format(sum_digit(123))); # "6"
    print("결과 : {}".format(sum_digit(777))); # "21"
