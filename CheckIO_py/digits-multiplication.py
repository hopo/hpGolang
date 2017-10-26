def checkio(number):
    s = str(number)
    r = 1
    for v in s:
        if not v == "0":
            r *= int(v)
    print(r)

if __name__ == '__main__':
    checkio(1000)
