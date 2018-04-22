
def reverse_str(string):
    ret = ""
    for s in string:
        ret = s + ret
    return ret

if __name__ == '__main__':
    ex1 = reverse_str("Zbcdefg") # "gfedcbZ"
    print(ex1)
    ex2 = reverse_str("dSdFlJ") # "JlFdSd"
    print(ex2)
