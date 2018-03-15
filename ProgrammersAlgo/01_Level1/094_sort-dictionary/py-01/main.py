def sort_dictionary(dic):
    tl = []
    for v in dic.items():
        s = v
        tl.append(s)
    tl.sort()
    return tl

print( sort_dictionary( {"김철수":78, "이하나":97, "정진원":88} ))
