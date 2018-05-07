
kor = 77
eng = 92
math = 82

if __name__ == '__main__':
    grade = ""
    average = (kor + eng + math) / 3
    if average >= 95:
        grade = "A+"
    elif average >= 90:
        grade = "A"
    elif average >= 85:
        grade = "B+"
    elif average >= 80:
        grade = "B"
    elif average >= 75:
        grade = "C+"
    elif average >= 70:
        grade = "C"
    elif average >= 65:
        grade = "D+"
    elif average >= 60:
        grade = "D"
    else :
        grade = "F"

    print("Your grade | {} : {}".format(grade, average))
