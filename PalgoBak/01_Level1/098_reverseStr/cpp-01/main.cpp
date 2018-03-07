#include <iostream>

using namespace std;

string reverseStr(string);

int main() {
	string ex1 = "Zbcdefg";     // "gfedcbZ"
	cout << reverseStr(ex1) << endl;
	
	string ex2 = "deSadFklJfg"; // "gfJlkFdaSed"
	cout << reverseStr(ex2) << endl;

	return 0;
}

string reverseStr(string s) {
	int lnth = strlen(s.c_str());

	int j;
	string ret;
	for(int i = 0; i < lnth; ++i) {
		j = lnth-1-i;
		ret.push_back(s[j]);
	}

	return ret;
}
