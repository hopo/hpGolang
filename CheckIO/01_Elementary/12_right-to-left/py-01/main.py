
def left_join(phrase: tuple) -> tuple:
    string = ""
    lnth = len(phrase)
    i = 0
    while (i < lnth):
        string += phrase[i]
        if not i == lnth - 1:
            string += ","
        i += 1
    ret = string.replace("right", "left")
    return ret

# ",",join(phrase)

if __name__ == '__main__':
    ex1 = left_join(("lorem","ipsum","dolor","sit","amet","consectetuer","adipiscing","elit","aenean","commodo","ligula","eget","dolor","aenean","massa","cum","sociis","natoque","penatibus","et","magnis","dis","parturient","montes","nascetur","ridiculus","mus","donec","quam","felis","ultricies","nec","pellentesque","eu","pretium","quis","sem","nulla","consequat","massa","quis",)) 
    print(ex1)
