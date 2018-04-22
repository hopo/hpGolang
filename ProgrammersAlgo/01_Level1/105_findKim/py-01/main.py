
def findKim(seoul): # (list) str
    kimIdx = -1
    for i in range(len(seoul)):
        if seoul[i] == "Kim":
            kimIdx = i
            break

    return "김서방은 {}에 있다".format(kimIdx)

if __name__ == '__main__':
    print(findKim(["Queen", "Tod", "Kim"])) # 2
