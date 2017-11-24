def left_join(phrase):
    s = ""
    l = len(phrase)
    i = 0
    while (i < l):
        s += phrase[i]
        if not i == l - 1:
            s += ","
        i += 1
    r = s.replace("right", "left")
    return r

# ",",join(phrase)

if __name__ == '__main__':
    ex1 = left_join(("lorem","ipsum","dolor","sit","amet","consectetuer","adipiscing","elit","aenean","commodo","ligula","eget","dolor","aenean","massa","cum","sociis","natoque","penatibus","et","magnis","dis","parturient","montes","nascetur","ridiculus","mus","donec","quam","felis","ultricies","nec","pellentesque","eu","pretium","quis","sem","nulla","consequat","massa","quis",)) 
    print(ex1)
