# https://py.checkio.org/mission/correct-sentence/

"""
For the input of your function will be given one sentence. You have to return its fixed copy in a way so it’s always starts with a capital letter and ends with a dot.
Pay attention to the fact that not all of the fixes is necessary. If a sentence already ends with a dot then adding another one will be a mistake.
Input: A string.
Output: A string.
Precondition: No leading and trailing spaces, text contains only spaces, a-z A-Z , and .
"""

def correct_sentence(text: str) -> str:
    return text


if __name__ == '__main__':
    print("Example:")
    print(correct_sentence("greetings, friends"))

    ex1 = correct_sentence("greetings, friends") == "Greetings, friends."
    print(ex1)
    ex2 = correct_sentence("Greetings, friends") == "Greetings, friends."
    print(ex2)
    ex3 = correct_sentence("Greetings, friends.") == "Greetings, friends."
    print(ex3)
    ex4 =  correct_sentence("hi") == "Hi"
    print(ex4)
