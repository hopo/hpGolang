from math import sqrt

def nextSqaure(n): # (int) int
    return "no" if sqrt(n) % 1 else int((sqrt(n)+1)**2)

if __name__ == '__main__':
    print("결과 : {}".format(nextSqaure(121))); # 144
