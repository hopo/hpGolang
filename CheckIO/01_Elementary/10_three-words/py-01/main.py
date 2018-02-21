def checkio(words):
    wl = words.split()
    ch = 0
    for v in wl:
        if v.isdigit():
            ch = 0
        else:
            ch += 1
            if ch == 3:
                return True
    return False

if __name__ == '__main__':
    ex1 = checkio("Hello 1 World hello")
    print(ex1)

    ex2 = checkio("Hell World Hello")
    print(ex2)
    
