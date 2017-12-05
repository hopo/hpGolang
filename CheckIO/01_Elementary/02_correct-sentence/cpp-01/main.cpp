#include <iostream>
#include <string>
#include <ctype.h>
#include <stdio.h>

using namespace std;

string correct_sentence(string text) { 
	for (int i = 0; i < sizeof(text); i++) {
		cout << text[i] << ":" << islower(text[i]) << "\n";
	}
	
	/*
    if text[0].islower:
        u = text[0].upper()
        text = text.replace(text[0], u, 1)
    if text[len(text)-1] != ".":
        text = text+"." 
		*/
    return "\n*end";
}


int main() {
    string ex1 = correct_sentence("greetings, friends"); // "Greetings, friends."
    //string ex2 = correct_sentence("Greetings, friends"); // "Greetings, friends."
    //string ex3 = correct_sentence("Greetings, friends."); // "Greetings, friends."
    //string ex4 = correct_sentence("hi"); // "Hi."

	cout << ex1;
	return 0;
}
