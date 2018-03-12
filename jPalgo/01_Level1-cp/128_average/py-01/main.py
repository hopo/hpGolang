def average(list):
    x = 0
    for v in list:
        x += v
    return x/len(list)

list = [5,3,4]
print("평균값 : {}".format(average(list)));
