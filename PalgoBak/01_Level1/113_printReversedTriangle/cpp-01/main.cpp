#include <iostream>

using namespace std;

string printReversedTriangle(int);

int main() {
	string ex1 = printReversedTriangle(3); // ***\n**\n*
	cout << ex1 << endl;

	return 0;
}

string printReversedTriangle(int n) {
	string ret;
	int i, j, k;

	for(i = 0; i < n; ++i) {
		k = n - i;	
		for(j = 0; j < k; ++j) ret += "*";
		if(k == 1) break;
		ret += '\n';
	}

	return ret;
}
