
def number_generator(x, n): # (int, int) list
    ret = []
    while n != 0:
        ret.insert(0, x*n)
        n -= 1
    return ret

if __name__ == '__main__':
    print(number_generator(3, 5)) # [3, 6, 9, 12, 15] 
