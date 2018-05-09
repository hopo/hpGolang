
def longest_palindrom(s):
    i = 0

    for n in range(len(s)-1, -1, -1): #for n in reversed(range(len(s))):
        if s[i] == s[n]:
            i += 1
        else:
            i = 0
    
    return i

if __name__ == '__main__':
    print(longest_palindrom("토마토맛토마토")) #7
    print(longest_palindrom("토마토맛있어")) #3
