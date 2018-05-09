
def is_pair(s):
    a = 0
    if (s[0] == '('):
        if (s[len(s)-1] == ')'):
            a += 1
            for ch in s:
                if ch == '(': a += 1
                if ch == ')': a -= 1 
    return a == 1 and True or False

if __name__ == '__main__':
    print( is_pair("(hello)()")) #True
    print( is_pair(")(")) #False
    print( is_pair(")()")) #False
    print( is_pair("()(")) #False
    print( is_pair("()()")) #True
    print( is_pair("()")) #True
    print( is_pair("((())")) #False

