#include <iostream>

using namespace std;

int main() {
	char aa[] = "hello";
	
	cout << "aa: " <<  aa << endl;
	cout << "aa+2: " << aa + 2 << endl; // llo
	cout << "*aa+3: " << char(*aa+3) << endl; // 104(h)+3 = 107(k)
	cout << "aa[1]+1: " << aa[1]+1 << endl; // 101(e, aa[1])+1 = 102

	return 0;
}
