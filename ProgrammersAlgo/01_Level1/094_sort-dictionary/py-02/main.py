
def sort_dictionary(dic): # (list) list
    ret = []
    for k in dic:
        ret.append((k, dic[k]))
    return sorted(ret)

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print( sort_dictionary( {"김철수":78, "이하나":97, "정진원":88} )) # [('김철수', 78), ('이하나', 97), ('정진원', 88)] 
