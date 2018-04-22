
def alpha_string46(string): # (str) bool

    ret = False
    if string.isdigit():
    	if (len(string) == 4) or (len(string) == 6):
    	    ret = True

    return ret

if __name__ == '__main__':
    print(alpha_string46("a234")) # False
    print(alpha_string46("1234")) # True
