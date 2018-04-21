import re

def find_message(text):
    """ Find a secret message """
    p = re.compile('[A-Z]')
    lst = p.findall(text)

    ret = ""
    for v in lst:
        ret += v
    """
    for v in text:
        if v == "x":
            ret += v
    "S".isupper() == true
    """
    return ret

if __name__ == '__main__':
    ex1 = find_message("How are you? Eh, ok. Low or Lower? Ohhh.") # "HELLO"
    print(ex1)

    ex2 = find_message("hello world!") # "", "Nothing"
    print(ex2)

    ex3 = find_message("HELLO WORLD!!!") # "HELLOWORLD", "Capitals"
    print(ex3)

