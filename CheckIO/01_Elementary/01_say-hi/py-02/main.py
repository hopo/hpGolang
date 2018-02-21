
def say_hi(name, age):
    ret = "Hi. My name is " + name + " and I'm " + str(age) + " years old"

    return ret

if __name__ == '__main__':
    name1 = "Alex"
    age1 = 32
    ex1 = say_hi(name1, age1) # "Hi. My name is Alex and I'm 32 years old", "First"
    print(ex1)

    name2 = "Frank"
    age2 = 68
    ex2 = say_hi(name2, age2) # "Hi. My name is Frank and I'm 68 years old", "Second"
    print(ex2)

