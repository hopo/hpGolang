#include <iostream>

using namespace std;

bool alpha_string46(string);

int main() {
	bool ex1 = alpha_string46("a234"); // 0, False, "string 'a'"
	cout << ex1 << endl;

	bool ex2 = alpha_string46("1234"); // 1, True, "all int"
	cout << ex2 << endl;

	return 0;
}

bool alpha_string46(string str) {
	int i = 0;
	while(1) {
		if(str[i] == 0) { break; }
		if('0' > str[i] || str[i] > '9') { return false; } 
		i++;
	}
	return true;
}
