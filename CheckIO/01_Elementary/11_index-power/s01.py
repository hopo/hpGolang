def index_power(array, n):
    if len(array) <= n:
        return -1
    else:
        return array[n] ** n

if __name__ == '__main__':
    ex1 = index_power([1, 2, 3, 4], 2) # 9, "Square"
    print(ex1)
    ex2 = index_power([1, 3, 10, 100], 3) # 1000000, "Cube"
    print(ex2)
    ex3 = index_power([0, 1], 0) # 1, "Zero power"
    print(ex3)
    ex4 = index_power([1, 2], 3) # -1, "IndexError"
    print(ex4)
