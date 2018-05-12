
def to_weirdCase(s):
    ret = ''
    flag = 0
    for x in s:
        ret += flag&1 and x or x.upper()
        if x == ' ':
            flag = 0
            continue
        flag = ~flag

    return ret

if __name__ == '__main__':
    ex1 = to_weirdCase("try hello world") # "TrY HeLlO WoRlD"
    print(ex1)

    ex2 = to_weirdCase("coding is thinking not typing") #"CoDiNg Is ThInKiNg NoT TyPiNg"
    print(ex2)
