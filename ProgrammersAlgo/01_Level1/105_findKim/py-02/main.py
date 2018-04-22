
def findKim(seoul): # (list) int
    kimIdx = seoul.index("Kim")
    return "김서방은 {}에 있다".format(kimIdx)

if __name__ == '__main__':
    print(findKim(["Queen", "Tod", "Kim"])) # 2
