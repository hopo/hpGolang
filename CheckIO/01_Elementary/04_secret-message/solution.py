import re

def find_message(text):
    """Find a secret message"""
    p = re.compile('[A-Z]')
    m = p.findall(text)

    ret = ""
    for v in m:
        ret += v
    #for v in text:
        #if v == "x":
            #ret += v
    #"S".isupper() == true
    return ret

if __name__ == '__main__':
    #These "asserts" using only for self-checking and not necessary for auto-testing
    r = find_message("How are you? Eh, ok. Low or Lower? Ohhh.")
    print(r)
    #assert find_message("hello world!") == "", "Nothing"
    #assert find_message("HELLO WORLD!!!") == "HELLOWORLD", "Capitals"
