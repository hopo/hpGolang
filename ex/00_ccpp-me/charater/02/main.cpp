#include <iostream>
#include <cstdio>

using namespace std;

void cprntf(char c) {
	printf("char : %c\n", c);
}

void caout(char* inpt) {
	cout << "inpt: "  << inpt << endl;
	cout << "inpt+1: "  << inpt+1 << endl;
	cout << "inpt[1]: " << inpt[1] << endl;
	cout << "inpt[1]+1: " << inpt[1]+1 << endl;
	cout << "size of inpt:" <<  sizeof(inpt);
}

int main() {
	char carr[] = "abcde";
	caout(carr);

	return 0;
}
