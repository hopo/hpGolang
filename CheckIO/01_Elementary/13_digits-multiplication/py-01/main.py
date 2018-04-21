
def checkio(number :int) -> int:
    string = str(number)
    r = 1
    for v in string:
        if not v == "0":
            r *= int(v)
    return r

if __name__ == '__main__':
    ex1 = checkio(1000) # 1
    print(ex1)
    ex2 = checkio(123405) #120
    print(ex2)
