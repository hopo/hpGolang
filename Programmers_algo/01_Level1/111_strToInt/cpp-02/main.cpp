#include <iostream>

using namespace std;

int strToInt(string);
int convctoi(char);

int main() {
	int ex1 = strToInt("-1234");	// -1234
	cout << ex1 << endl;
	int ex2 = strToInt("392");	// 392
	cout << ex2 << endl;

	return 0;
}

int strToInt(string s) {
	int lnth = strlen(s.c_str());
	int ret, i, j, p = 1;
	for(i = 0; i < lnth; ++i) {
		j = lnth-1-i;
		if('0' <= s[j] && s[j] <= '9') { 
			ret += convctoi(s[j])*p;
		}
		p *= 10;
	}

	return (s[0] == '-')? ret *= -1 : ret;
}

int convctoi(char c) {
	return (int)(c-48);	
}
