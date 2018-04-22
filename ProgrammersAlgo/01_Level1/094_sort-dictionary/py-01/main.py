
def sort_dictionary(dic): # (list) list
    tl = []
    for v in dic.items():
        s = v
        tl.append(s)
    tl.sort()

    return tl

if __name__ == '__main__':
    print( sort_dictionary( {"김철수":78, "이하나":97, "정진원":88} )) # [('김철수', 78), ('이하나', 97), ('정진원', 88)]
