
def alpha_string46(s): # (str) bool
    return s.isdigit() and len(s) in [4, 6]

if __name__ == '__main__':
    print(alpha_string46("a234")) # False
    print(alpha_string46("1234")) # True
