#include <iostream>

using namespace std;

string water_melon(int);

int main() {
	string ex1 = water_melon(3);	// "WMW"
	cout << ex1 << endl;

	string ex2 = water_melon(4);	// "WMWM"
	cout << ex2 << endl;

	return 0;
}

string water_melon(int num) {
	string ret;
	for(int i = 1; i < num+1; ++i) {
		(i%2 == 1)? (ret += "W") : (ret += "M");
	}

	return ret;
}
