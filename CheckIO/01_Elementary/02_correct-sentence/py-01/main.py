
def correct_sentence(text: str) -> str:
    # 타입 str의 일부 탐색
    if text[0].islower:
        u = text[0].upper()
        text = text.replace(text[0], u, 1) # str의 일부 바꾸기. replace()

    if text[len(text)-1] != ".":
        text = text + "." 

    return text

if __name__ == '__main__':
    ex1 = correct_sentence("greetings, friends") # "Greetings, friends."
    print(ex1)
    ex2 = correct_sentence("Greetings, friends") # "Greetings, friends."
    print(ex2)
    ex3 = correct_sentence("Greetings, friends.") # "Greetings, friends."
    print(ex3)
    ex4 =  correct_sentence("hi") # "Hi."
    print(ex4)
