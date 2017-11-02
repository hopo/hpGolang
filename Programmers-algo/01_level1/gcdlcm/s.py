import math
def gcdlcm(a, b):
    x = math.gcd(a, b)
    y = (a * b) / x
    answer = [x,int(y)]

    return answer


ex = gcdlcm(4,12)
print(ex)
