#include <iostream>

using namespace std;

string printTriangle(int);
string strMulti(string, int);

int main() {
	string ex1 = printTriangle(3); // *\n**\n***
	cout << ex1 << endl;

	return 0;
}

string printTriangle(int num) {
	string ret;
	for(int i = 1; i < num+1; ++i) {
		ret += strMulti("*", i);
		if(i != num) { ret += '\n'; } 
	}
	return ret;
}

string strMulti(string str, int n) {
	string ret;
	for(int i = 0; i < n; ++i) { ret += str; }

	return ret;
}
