
def sumMatrix(A, B): # (list, list) list
    # make multi matrix for return
    answer = [[0 for n in range(len(A[0]))] for n in range(len(A))]

    for i in range(len(A)):
        for j in range(len(A[0])):
        	answer[i][j] = A[i][j]+B[i][j]
    return answer

# 아래는 테스트로 출력해 보기 위한 코드입니다.
if __name__ == '__main__':
    print(sumMatrix([[1,2], [2,3]], [[3,4],[5,6]])) # [[4, 6], [[7, 9]]
