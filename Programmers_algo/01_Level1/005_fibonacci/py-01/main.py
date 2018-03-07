def fibonacci(num):
    a, b = 1, 0
    while num != 0:
        a, b = a+b, a
        num -= 1
    return b

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(fibonacci(3))
