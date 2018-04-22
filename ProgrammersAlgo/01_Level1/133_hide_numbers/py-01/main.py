
def hide_numbers(s): # (str) str
    ln = len(s)
    ln4 = ln-4
    front = '*' * ln4
    end = s[ln4:ln]
    return front + end

if __name__ == '__main__':
    print("결과 : " + hide_numbers('01033334444')); # "*******4444"
