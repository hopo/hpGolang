def nextSqaure(n):
    sr = n**(1/2)
    return int(sr)**2 == n and (sr+1)**2 or 'no'

print("결과 : {}".format(nextSqaure(121)));
