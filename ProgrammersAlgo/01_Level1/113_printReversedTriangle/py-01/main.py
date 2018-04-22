
def print_reversed_triangle(num):
    ret = ""
    while num != 0:
        ret += num * "*"
        ret += "\n"
        num -= 1

    return ret

if __name__ == '__main__':
    ex1 = print_reversed_triangle(3) # "***\n**\n*\n"
    print(ex1)
