def number_generator(x, n):
    ret = []
    while n != 0:
        ret.insert(0, x*n)
        n -= 1
    return ret

print(number_generator(3,5))
