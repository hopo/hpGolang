def say_hi(name, age):
    return "Hi. My name is " + name + " and I'm " + str(age) + " years old"

if __name__ == '__main__':
    ex1 = say_hi("Alex", 32) == "Hi. My name is Alex and I'm 32 years old", "First"
    print(ex1)
    ex2 = say_hi("Frank", 68) == "Hi. My name is Frank and I'm 68 years old", "Second"
    print(ex2)

