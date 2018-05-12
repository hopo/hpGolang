
def adder(a, b):
    if a > b: a, b = b, a
    return sum(range(a, b+1))

if __name__ == '__main__':
    ex1 = adder(3, 5) # 12, 3+4+5
    print(ex1)
