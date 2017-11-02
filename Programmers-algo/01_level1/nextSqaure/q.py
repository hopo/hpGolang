# https://programmers.co.kr/learn/challenge_codes/119

def nextSqure(n):
    sr = n**(1/2)
    return int(sr)**2 == n and (sr+1)**2 or 'no'

# 아래는 테스트로 출력해 보기 위한 코드입니다.
print("결과 : {}".format(nextSqure(121)));
