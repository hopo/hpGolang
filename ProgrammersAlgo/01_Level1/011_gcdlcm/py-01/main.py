
def gcdlcm(a, b):
    # max and min
    if a > b: a, b = b, a
    mn, mx = a, b

    while mn > 0:
        mx, mn = mn, mx%mn

    return [mx, int((a*b)/mx)]

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print(gcdlcm(3,12))
