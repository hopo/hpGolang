
def checkio(number):
    if number % 15 == 0:
        return "Fizz Buzz"
    elif number % 3 == 0:
        return "Fizz"
    elif number % 5 == 0:
        return "Buzz"

    return str(number) # int -> str casting

if __name__ == '__main__':
    ex1 = checkio(15) # "Fizz Buzz", "15 is divisible by 3 and 5"
    print(ex1)
    ex2 = checkio(6) # "Fizz", "6 is divisible by 3"
    print(ex2)
    ex3 = checkio(5) # "Buzz", "5 is divisible by 5"
    print(ex3)
    ex4 = checkio(7) # "7", "7 is not divisible by 3 or 5"
    print(ex4)
