
def average(mu_list): # (list) float
    total = 0
    for v in my_list:
        total += v
    return total/len(my_list)

if __name__ == '__main__':
    my_list = [5, 3, 4]
    print("평균값 : {}".format(average(my_list))); # 4
