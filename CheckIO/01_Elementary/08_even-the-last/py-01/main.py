
def checkio(array):
    if array:
        sum = 0
        for i in range(len(array)):
            if i%2 == 0:
                sum += array[i]
        return sum * array[-1]
    else:
        return 0

if __name__ == '__main__':
    ex1 = checkio([0, 1, 2, 3, 4, 5]) # 30, "(0+2+4)*5=30"
    print(ex1)

    ex2 = checkio([1, 3, 5]) # 30, "(1+5)*5=30"
    print(ex2)

    ex3 = checkio([6]) # 36, "(6)*6=36"
    print(ex3)

    ex4 = checkio([]) # 0, "An empty array = 0"
    print(ex4)
