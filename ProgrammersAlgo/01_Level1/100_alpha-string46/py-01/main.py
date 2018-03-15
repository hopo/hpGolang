def alpha_string46(s):
    if s.isdigit():
    	if len(s) == 4 or len(s) == 6:
    		return True
    return False

print( alpha_string46("a234") )
print( alpha_string46("1234") )
