#include <iostream>

using namespace std;

bool numPY(string);

int main() {
	string ex1 = "pPoooyY"; // 1, true
	cout << numPY(ex1) << endl;

	string ex2 = "Pyy";     // 0, false
	cout << numPY(ex2) << endl;

	return -1;
}

bool numPY(string s) {
	int ppp, yyy;
	for(auto& v : s) {
		if(v == 'p' || v == 'P') { ppp++; }
		if(v == 'y' || v == 'Y') { yyy++; }
	}
	return ppp == yyy;
}
