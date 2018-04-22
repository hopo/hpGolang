
def printTriangle(num): # (int) str
    ret= ""
    i = 1
    while (i < num+1):
        ret += "*" * i
        i += 1
        ret += "\n"

    return ret

if __name__ == '__main__':
    print( printTriangle(3) ) # "*\n**\n***\n"
