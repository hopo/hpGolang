
def strange_sort(strings, num): # (list, int) list
    for i, v in enumerate(strings):
        for j, w in enumerate(strings):
            if v[num] < w[num]:
                strings[i], strings[j] = strings[j], strings[i]

    return strings

if __name__ == '__main__':
    print(strange_sort(["sun", "bed", "car"], 1)) # ['car', 'bed', 'sun'] 

