def printTriangle(num):
    s = ""
    i = 1
    while (i < num+1):
        s += "*" * i
        i += 1
        s += "\n"
    return s

print( printTriangle(3) )
