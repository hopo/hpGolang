
def divisible(iarr, num):
    ret = []

    for v in iarr:
        if not v%num:
            ret.append(v)

    return ret

if __name__ == '__main__':
    ex1 = divisible([5, 9, 7, 10], 5) # [5, 10]
    print(ex1)
