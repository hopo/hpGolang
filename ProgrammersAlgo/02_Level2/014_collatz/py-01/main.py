
def collats(num):
    count = 0
    while num != 1:
        if num&1:
            num = num*3 + 1
        else:
            num //= 2
        count += 1
    
    return count

if __name__ == '__main__':
    ex1 = collats(6)  # 8, 3->10->5->16->8->4->2->1
    print(ex1)
    
    ex2 = collats(11) #14, 34->17->52->26->13->40->20->10->5->16->8->4->2->1
    print(ex2)
