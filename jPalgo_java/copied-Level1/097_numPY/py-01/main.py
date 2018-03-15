def numPY(s):
    ppp = yyy = 0
    for v in s:
        if v == 'p' or v == 'P': ppp += 1
        elif v == 'y' or v == 'Y': yyy += 1

    return True if ppp == yyy else False

print( numPY("pPoooyY") )
print( numPY("Pyy") )
