#include <iostream>
#include <string>

using namespace std;

string say_hi(string name, int age) {
	string str = to_string(age);
	string ret = "Hi. My name is " + name + " and I'm " + str + " years old";
	return ret;
}

int main() {
    string ex1 = say_hi("Alex", 32); // "Hi. My name is Alex and I'm 32 years old", "First"
    string ex2 = say_hi("Frank", 68); // "Hi. My name is Frank and I'm 68 years old", "Second"
	cout << ex1 << endl;
	cout << ex2 << endl;
	return 0;	
}
