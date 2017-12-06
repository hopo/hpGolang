#include <iostream>
#include <string>

using namespace std;

void myntcs() {
	char myntcs[] = "some text";
	string mystring = myntcs;	// convert c-string
	cout << mystring << "/1/\n";	// printed as a library string
	cout << mystring.c_str() << "/2/\n";	// printed as a c-string
}

int main() {
	myntcs();
	return 0;
}
