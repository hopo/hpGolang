
"""
def gcdlcm(a, b):
    gcd = two(b)
    lcm = three(a)
    answer = [gcd, lcm]

    return answer

"""
def two(num):
    if num % 2 == 1:
        return num
    return num % 2

def three(num):
    if num % 2 == 0:
        return num
    return num / 3

def pfac(n):
    x = two(n)
    return x

print(pfac(7))
