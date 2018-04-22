
def evenOrOdd(num): # (int) str
    if num%2 == 0: return 'Even'
    return 'Odd'

if __name__ == '__main__':
    print("결과 : " + evenOrOdd(3)) # "Odd"
    print("결과 : " + evenOrOdd(2)) # "Even"

