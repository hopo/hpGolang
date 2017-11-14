def checkio(array):
    if array:
        sum = 0
        for n in range(len(array)):
            if n % 2 == 0:
                sum = sum+array[n]
        return sum*array[-1]
    else:
        return 0

if __name__ == '__main__':
    assert checkio([0, 1, 2, 3, 4, 5]) == 30, "(0+2+4)*5=30"
    assert checkio([1, 3, 5]) == 30, "(1+5)*5=30"
    assert checkio([6]) == 36, "(6)*6=36"
    assert checkio([]) == 0, "An empty array = 0"
