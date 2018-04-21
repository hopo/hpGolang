
def checkio(*args) -> int:
    nums  = []
    if args == (): return 0

    for v in args: nums.append(v)

    nums.sort()
    mn = nums[0]
    mx = nums[len(nums)-1]
    return mx - mn

if __name__ == '__main__':
    ex1 = checkio(2, 3, 1)  # 2, (1, 2, 3), 3-1
    print(ex1)

    ex2 = checkio() # 0
    print(ex2)
