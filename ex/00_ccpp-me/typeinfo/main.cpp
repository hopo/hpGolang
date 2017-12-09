#include <iostream>
#include <typeinfo>
#include <string>
#include <vector>

using namespace std;

class someclass {};

int main() {
	int aa = -2;
	char cc = 'h';
	string mm = "lets";
	int xx;
	string yy;
	someclass zz;

	cout << typeid(vector<int>).name() << "!" << '\n';

	cout << typeid(aa).name() << '\n';
	cout << typeid(cc).name() << '\n';
	cout << typeid(mm).name() << '\n';
	cout << typeid(xx).name() << '\n';
	cout << typeid(yy).name() << '\n';
	cout << typeid(zz).name() << '\n';

	return 0;
}
