
def nextSqaure(num): # (int) int
    sr = num**(1/2) # sqrt(num, 0.5)
    return int(sr)**2 == num and int((sr+1)**2) or "no"

if __name__ == '__main__':
    print("결과 : {}".format(nextSqaure(121))); # 144
