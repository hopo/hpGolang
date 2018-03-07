def nextSqaure(n):
    from math import sqrt
    return "no" if sqrt(n) % 1 else (sqrt(n)+1)**2

print("결과 : {}".format(nextSqaure(121)));
