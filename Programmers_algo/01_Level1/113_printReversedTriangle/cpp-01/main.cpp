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
	int i, j;
	/*
	for i := 0; i < n; i++ {
		q := n - i
		r := hpkg.Smulti("*", q)
		switch q {
		default:
			ret += fmt.Sprintln(r)
		case 1:
			ret += fmt.Sprint(r)
		}
	}
	*/
	return "-1";
}
