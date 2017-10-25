def checkio(*args):
    l = [] 
    for v in args:
       l.append(v)
    l.sort()
    a = l[0]
    b = l[len(l)-1]
    
    return b - a

ret = checkio(2, 3, 1)

print(ret)
