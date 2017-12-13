#include <iostream>

using namespace std;

int main() {
	int i = 0, j = 0, n = 10;	

	do {
		i++;
		cout << i << ", ";
	}
	while(i < n);
	cout << '\n';

	do {
		++j;
		cout << j << ", ";
	}
	while(j < n);

	return 0;
}
