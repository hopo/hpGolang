def checkio(numbers_array):
    i, j = 0, 1
    while (j < l):
        if abs(arr[i]) < abs(arr[j]):
            arr[i], arr[j] = arr[j], arr[i]
        i += 1
        j += 1
    print(numbers_array)
    return 0

def len(): return len(arr)
def swap(a, b): a, b = b, a
def less(a, b): return abs(a) < abs(b)

if __name__ == '__main__' :
    checkio((-3, 2, 0))
