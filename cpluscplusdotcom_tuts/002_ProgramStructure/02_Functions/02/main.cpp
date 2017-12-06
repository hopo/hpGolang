#include <iostream>

using namespace std;

int substraction (int a, int b)
{
	int r;
	r = a-b;
	return r;
}

int main()
{
	int x = 5, y = 3, z;
	z = substraction (7,2);
	cout << "The first result is " << z << '\n';
	cout << "The second result is " << substraction (7, 2) << '\n';
	cout << "The third result is " << substraction (x, y) << '\n';
	z = 4 + substraction (x, y);
	cout << "The fourth result is " << z << '\n';
}
