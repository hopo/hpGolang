// ing

#include <iostream>

using namespace std;

int nextSqaure(int);

int main() {
	int ex1 = nextSqaure(121); // 144, 12*12
	cout << ex1 << endl;
	// int ex2 = nextSqaure(4);   // 9, 3*3
	// cout << ex2 << endl;
	// int ex3 = nextSqaure(5);   // 0, "no"
	// cout << ex3 << endl;

	return 0;
}

int nextSqaure(int n) {
	cout << "n: " << n << endl;
	/*
	x := math.Sqrt(float64(n))
	f := math.Pow(x, 2)
	if f == float64(n) {
		rslt = int(math.Pow(x+1, 2))
	} else {
		rslt = -1
	}
	*/

	return -1;
}
