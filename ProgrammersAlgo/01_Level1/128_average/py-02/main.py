
def average(array): # (list) float
    return sum(array)/len(array)

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    list = [5, 3, 4]
    print("평균값 : {}".format(average(list))); # 4.0
