def findKim(seoul):
    kimIdx = 0
    for i in range(len(seoul)):
        if seoul[i] == "Kim":
            kimIdx = i
            break
            
    return "김서방은 {}에 있다".format(kimIdx)

print(findKim(["Queen", "Tod", "Kim"]))
