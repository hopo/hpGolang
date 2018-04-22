
def no_continuous(string): #(str) list
    ret = []

    for i in range(len(string)):
        if string[i-1] != string[i]:
            ret.append(string[i])

    return ret

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print(no_continuous("133303")) # ['1', '3', '0', '3']

