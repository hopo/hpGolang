
def get_minmax_string(string):
    sarr = string.split(" ")
    iarr = []
    for s in sarr:
        iarr.append(int(s))

    return [min(iarr), max(iarr)]

if __name__ == '__main__':
    ex1 = get_minmax_string("1 2 3 4") # [1, 4]
    print(ex1)
