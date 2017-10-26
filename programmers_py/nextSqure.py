def nextSqure(n):
    sr = n**(1/2)
    return int(sr)**2 == n and (sr+1)**2 or 'no'
"""
def nextSqure(n):
from math import sqrt
    return "no" if sqrt(n) % 1 else (sqrt(n)+1)**2
"""
print("결과 : {}".format(nextSqure(121)));
