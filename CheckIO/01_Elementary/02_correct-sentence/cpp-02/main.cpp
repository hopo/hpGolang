#include <iostream>
#include <string>
#include <cstring>

using namespace std;

string correct_sentence(char* text) {
	string s = text; // convert c-string to string
	size_t l = strlen(text); 

	// text[0] change upper case.
	if (islower(s[0])) { s[0] = char(s[0]-32); }
	// text end "."? 
	if (s[l-1] != '.') { s.insert(s.end(), '.'); }

	return s;
}

int main() {
	char ex1[] = "greeting, friends";
	cout << correct_sentence(ex1) << endl; // "Greeting, friends."

	char ex2[] = "Greeting, friends"; // "Greeting, friends."
	cout << correct_sentence(ex2) << endl;

	char ex3[] = "Greeting, friends."; // "Greeting, friends."
	cout << correct_sentence(ex3) << endl;

	char ex4[] = "hi"; // "Hi."
	cout << correct_sentence(ex4) << endl;

	return 0;
}
