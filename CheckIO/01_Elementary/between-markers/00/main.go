// ...ing
package main

import (
	"fmt"
)

func main() {
	ex1 := between_markers("What is >apple<", ">", "<") // "apple", "One sym"
	fmt.Println(ex1)
	/*
		ex2 := between_markers("<head><title>My new site</title></head>", "<title>", "</title>") // "My new site", "HTML"
		fmt.Println(ex2)
		ex3 := between_markers("No[/b] hi", "[b]", "[/b]") // "No", 'No opened'
		fmt.Println(ex3)
		ex4 := between_markers("No [b]hi", "[b]", "[/b]") // "hi", 'No close'
		fmt.Println(ex4)
		ex5 := between_markers("No hi", "[b]", "[/b]") // "No hi", 'No markers at all'
		fmt.Println(ex5)
		ex6 := between_markers("No <hi>", ">", "<") // "", 'Wrong direction'
		fmt.Println(ex6)
	*/
}

func between_markers(text, begin, end string) string {
	bsl := []byte(text)
	bbeg := []byte(begin)
	bend := []byte(end)
	var oi, ci int
	for i, _ := range bsl {
		if bsl[i] == bbeg[0] {
			oi = i
		}
		if bsl[i] == bend[0] {
			ci = i
		}
	}
	var ret []byte
	ret = append(ret, bsl[oi+1:ci]...)
	return string(ret)
}

/*
# https://py.checkio.org/mission/between-markers/

"""
You are given a string and two markers (the initial and final). You have to find a substring enclosed between these two markers. But there are a few important conditions:
The initial and final markers are always different.
The text must be found only between the first instances of the marker.
If there is no initial marker then the beginning should be considered as the beginning of a string.
If there is no final marker then the ending should be considered as the ending of a string.
If the initial and final markers are missing then simply return the whole string.
If the final marker is standing in front of the initial one then return an empty string.
Input: Three arguments. All of them are strings. The second and third arguments are the initial and final markers.
Output: A string.
Precondition: can't be more than one marker
"""

def between_markers(text: str, begin: str, end: str) -> str:
    """
        returns substring between two given markers
    """
    # your code here
    return ''


if __name__ == '__main__':
    print('Example:')
    print(between_markers('What is >apple<', '>', '<'))

    # These "asserts" are used for self-checking and not for testing
    assert between_markers('What is >apple<', '>', '<') == "apple", "One sym"
    assert between_markers("<head><title>My new site</title></head>",
                           "<title>", "</title>") == "My new site", "HTML"
    assert between_markers('No[/b] hi', '[b]', '[/b]') == 'No', 'No opened'
    assert between_markers('No [b]hi', '[b]', '[/b]') == 'hi', 'No close'
    assert between_markers('No hi', '[b]', '[/b]') == 'No hi', 'No markers at all'
    assert between_markers('No <hi>', '>', '<') == '', 'Wrong direction'
    print('Wow, you are doing pretty good. Time to check it!')

*/
