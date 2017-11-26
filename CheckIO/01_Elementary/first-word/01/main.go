package main

import (
	"fmt"
)

func main() {
	ex1 := first_word("Hello world") // "Hello"
	fmt.Println(ex1)
	ex2 := first_word(" a word ") // "a"
	fmt.Println(ex2)
	ex3 := first_word("don't touch it") // "don't"
	fmt.Println(ex3)
	ex4 := first_word("greetings, friends") // "greetings"
	fmt.Println(ex4)
	ex5 := first_word("... and so on ...") // "and"
	fmt.Println(ex5)
	ex6 := first_word("hi") // "hi"
	fmt.Println(ex6)
}

func first_word(text string) string {
	return "!"
}

/*
# https://py.checkio.org/mission/first-word/

"""
You are given a string where you have to find its first word.
When solving a task pay attention to the following points:
There can be dots and commas in a string.
A string can start with a letter or, for example, a dot or space.
A word can contain an apostrophe and it's a part of a word.
The whole text can be represented with one word and that's it.
Input: A string.
Output: A string.
Precondition: the text can contain a-z A-Z , . '
"""

def first_word(text: str) -> str:
    """
        returns the first word in a given text.
    """
    # your code here
    return text[0:2]


if __name__ == '__main__':
    print("Example:")
    print(first_word("Hello world"))

    # These "asserts" are used for self-checking and not for an auto-testing
    assert first_word("Hello world") == "Hello"
    assert first_word(" a word ") == "a"
    assert first_word("don't touch it") == "don't"
    assert first_word("greetings, friends") == "greetings"
    assert first_word("... and so on ...") == "and"
    assert first_word("hi") == "hi"
    print("Coding complete? Click 'Check' to earn cool rewards!")
*/
