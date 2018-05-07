def num_of_prime(num):
    pm_count = 0

    for n in range(2, num+1):
        if is_prime(n):
            pm_count += 1

    return pm_count


def is_prime(num):
    isp = True; 

    for n in range(2, num//2+1):
        if num%n == 0:
            isp = False
            break

    return isp


if __name__ == '__main__':
    ex1 = num_of_prime(10) # 4, [2 3 5 7]
    print(ex1)

    ex2 = num_of_prime(5) # 3, [2 3 5]
    print(ex2)

