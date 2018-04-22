
def water_melon(num): # (int) str
    ret = ""
    for i in range(num):
    	ret += ~i&1 and "수" or "박"

    return ret

# 실행을 위한 테스트코드입니다.
if __name__ == '__main__':
    print("n이 3인 경우: " + water_melon(3)); # "수박수"
    print("n이 4인 경우: " + water_melon(4)); # "수박수박"
