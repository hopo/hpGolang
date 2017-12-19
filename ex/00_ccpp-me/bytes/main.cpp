#include <cstdio>
#include <iostream>

using namespace std;

void foo() {
	int x = 120;
	printf("01: %c\n", x);	

	char a[] = "ab";
	printf("02: %c\n", a[0]);
	std::cout << a[1];

	char b[] = "xyz";
	cout << "03: " <<  b[1];
	b[1] = 99;	// 99 or 'c'
	cout << "03': " <<  b[1];
}

int main() {
	foo();
	return 0;
}
