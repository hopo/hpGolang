#include <iostream>

using namespace std;

class MyClass
{
	public:
		int x;
		MyClass(int val) :x(val) {}
		int get() { return x; }
};

int main() {
	const MyClass foo(10);
	// foo.x = 20;	// not valid: x cannot be modified
	cout << foo.x << '\n';
	return 0;
}
