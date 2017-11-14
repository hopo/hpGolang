def checkio(numbers_array):
    arr = []
    for v in numbers_array:
        arr.append(v)

    l = len(arr)
    i = 1
    while (i < l):
        j = 0
        while (j < l):
            if abs(arr[i]) < abs(arr[j]):
                arr[i], arr[j] = arr[j], arr[i]
            j += 1
        i += 1
    return arr

if __name__ == '__main__' :
   a = checkio((1,2,3,0))
   print(a)
