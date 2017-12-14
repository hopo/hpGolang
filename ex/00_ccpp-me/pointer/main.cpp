#include <iostream>
#include <typeinfo>

using namespace std;

int main() {
	string str = "hello";

	string* pstr = &str;

	string& astr = str;

	cout << "num: " << str << endl; 
	cout << " 1- " << str << ",  " << *pstr << endl;	// hello, hello
	cout << " 2- " << &str << ",  " << pstr << endl;	// adress, adress
	cout << " 3- " << typeid(astr).name() <<  ", " << astr << endl;	// blahblahtype, hello
	cout << " 4- " << sizeof(str) << ",  " << sizeof(&str) << endl;	// 24, 4
	cout << " 5- " << str.size() << ", " << pstr->size() << endl;	// 5, 5

	return 0;	
}
