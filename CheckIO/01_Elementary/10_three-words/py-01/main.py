
def checkio(words) -> bool:
    word = words.split()
    chk = 0
    for v in word:
        if v.isdigit():
            chk = 0
        else:
            chk += 1
            if chk == 3:
                return True
    return False

if __name__ == '__main__':
    ex1 = checkio("Hello 1 World hello") # False
    print(ex1)

    ex2 = checkio("Hell World Hello") # True
    print(ex2)
    
