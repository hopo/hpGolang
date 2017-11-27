def checkio(*args):
    l  = []
    if args == ():
        return 0
    for v in args:
       l.append(v)
    l.sort()
    a = l[0]
    b = l[len(l)-1]
    return b - a


if __name__ == '__main__':
    ex1 = checkio(2, 3, 1)
    print(ex1)
    ex2 = checkio()
    print(ex2)
