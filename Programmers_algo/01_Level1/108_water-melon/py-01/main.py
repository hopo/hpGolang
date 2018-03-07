def water_melon(n):
    ret = ""
    for i in range(n):
    	ret += ~i&1 and "수" or "박"
    return ret


# 실행을 위한 테스트코드입니다.
print("n이 3인 경우: " + water_melon(3));
print("n이 4인 경우: " + water_melon(4));
