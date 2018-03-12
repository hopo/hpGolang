def strange_sort(strings, n):
    for i, v in enumerate(strings):
        for j, w in enumerate(strings):
            if v[n] < w[n]:
                strings[i], strings[j] = strings[j], strings[i]
    return strings

print( strange_sort(["sun", "bed", "car"], 1) )
