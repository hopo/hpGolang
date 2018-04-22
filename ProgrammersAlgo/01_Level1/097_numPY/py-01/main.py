
def numPY(string): # (str) bool
    ppp = yyy = 0
    for v in string:
        if v == 'p' or v == 'P': ppp += 1
        elif v == 'y' or v == 'Y': yyy += 1

    #return True if ppp == yyy else False
    return ppp == yyy and True or False

if __name__ == '__main__':
    print(numPY("pPoooyY")) # True
    print(numPY("Pyy")) # False
