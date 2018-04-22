
def string_middle(word): # (str) str
    return len(word)&1 and word[len(word)//2] or word[len(word)//2-1:len(word)//2+1]

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print(string_middle("power")) # "w"
    print(string_middle("main")) # "ai"
