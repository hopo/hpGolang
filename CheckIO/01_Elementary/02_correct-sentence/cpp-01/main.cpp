// ...ing
#include <iostream>
#include <string>
#include <ctype.h>
#include <stdio.h>

using namespace std;

string correct_sentence(string text) { 
	string ret = text;
	if (islower(ret[0])) {
		ret[0] = ret[0]-32; 
	}
	
	int l = sizeof(ret);

	cout << l;
    return "\n***";
}

int main() {
    //string ex1 = correct_sentence("greetings, friends"); // "Greetings, friends."
    //string ex2 = correct_sentence("Greetings, friends"); // "Greetings, friends."
    string ex3 = correct_sentence("Greetings, friends."); // "Greetings, friends."
    //string ex4 = correct_sentence("hi"); // "Hi."

	//cout << ex1 << endl;
	//cout << ex2 << endl;
	cout << ex3 ;
	//cout << ex4 << endl;
}

