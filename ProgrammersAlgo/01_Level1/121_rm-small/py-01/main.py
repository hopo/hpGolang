
def rm_small(mylist): # (list) list
    mylist.remove(min(mylist))
    return mylist

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    my_list = [4, 3, 2, 1]
    print("결과 {} ".format(rm_small(my_list))) # [4, 3, 2]
