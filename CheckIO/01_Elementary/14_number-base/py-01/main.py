
ALP = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

def checkio(str_number:str, radix:int) -> int:
    l = len(str_number)
    il = []
    for i in range(l):
        if str_number[i] in ALP:
            e = ALP.index(str_number[i]) + 10
        else:
            e = int(str_number[i])
        if e >= radix:
            return -1
        else:
            r = e * radix**(l-(i+1))
            il.append(r)
    return sum(il)

if __name__ == '__main__':
    ex1 = checkio("AF", 16) # 175
    print(ex1)
    ex2 = checkio("101", 5) # 26
    print(ex2)
