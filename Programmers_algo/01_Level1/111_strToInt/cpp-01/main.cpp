#include <iostream>
#include <algorithm>

using namespace std;

int strToInt(string);

int main() {
	int ex1 = strToInt("-1234");	// -1234
	cout << ex1 << endl;

	return 0;
}

int strToInt(string s) {
	return atoi(s.c_str());	
}
