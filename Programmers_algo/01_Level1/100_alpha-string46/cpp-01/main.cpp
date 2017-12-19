#include <iostream>

using namespace std;

bool alpha_string46(string);

int main() {
	bool ex1 = alpha_string46("a234"); // False, "string 'a'"
	cout << ex1 << endl;

	bool ex2 = alpha_string46("1234"); // True, "all int"
	cout << ex2 << endl;

	return 0;
}

bool alpha_string46(string str) {
	cout << str << endl;
	return false;
}
