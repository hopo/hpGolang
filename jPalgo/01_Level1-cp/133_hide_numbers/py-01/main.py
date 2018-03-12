def hide_numbers(s):
    ln = len(s)
    ln4 = ln-4
    front = '*' * ln4
    end = s[ln4:ln]
    return front + end

print("결과 : " + hide_numbers('01033334444'));
