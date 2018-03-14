#include <iostream>
#include <string>

using namespace std;

string correct_sentence(char* text) {
	// text[0] change upper case.
	char t = text[0];
	if (t > 96 && t < 123) { /* islower(t) */
		text[0] = char(t-32);
	}
	if (t > 64 && t < 91) { /* isupper(t) */
	}

	// end char[] is '.'?
	int l;
	for (l = 0; ;l++) {
		if (text[l] == 0 ) {
			break;
		}
	} /* cstring::strlen */

	string s = text;
	if (s[l-1] != '.') {
		s.insert(s.end(), '.');
	}

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
