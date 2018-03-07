def alpha_string46(s):
    return s.isdigit() and len(s) in [4, 6]

print( alpha_string46("a234") )
print( alpha_string46("1234") )
