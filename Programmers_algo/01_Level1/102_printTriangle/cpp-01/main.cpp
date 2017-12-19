#include <iostream>

using namespace std;

string printTriangle(int);

int main() {
	string ex1 = printTriangle(3); // *\n**\n***
	cout << ex1 << endl;

	return 0;
}

string printTriangle(int num) {
	string ret;
	for(int i = 1; i < num+1; i++) {
		for(int j = 0; j < i; j++) { ret.push_back('*'); }
		ret.push_back('\n');
	}
	return ret;
}
