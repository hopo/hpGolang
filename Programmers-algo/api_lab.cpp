// *********************
// *** For cpp api by hp
// *********************


// reference
#include <iostream>
#include <vector>

using namespace std;


// method list in here
void vctrprt();
vector<int> denom();
vector<int> fibonacci();


// *** vctrprt()
// p : int
// r : void
// vector int print like array type
void vctrprt(vector<int> vctr) {
	cout << "{";
	for (int i = 0; i < vctr.size(); i++) { // vector::size()
		if (i == vctr.size()-1) {
			cout << vctr[i];
			break;
		}
		cout << vctr[i] << ", ";
	}
	cout << "}" << endl;
}


// *** denom()
// p : int
// r : vector<int> 
// denominator 1 ~ n
vector<int> denom(int d) {
	vector<int> box;
	for (int i = 1; i < d+1; i++) {
		if (d % i == 0) {
			box.push_back(i); // vector::push_back()
		}
	}
	return box;
}

// *** fibonacci()
// p : int 
// r : vector<int>
// size = n, fibonacci vector int
// vector reference
vector<int> fibonacci(int n) {
	long long a, b, c;
	a = 0;
	b = 1;
	c = a + b;
	vector<int> box;
	for (int i = 0; i < n ; i++) {
		a = b;
		b = c;
		box.push_back(b);
		c = a + b;
	}
	return box;
}


// !
// main method for check
int main() {
	auto d = denom(12);
	vctrprt(d);

	auto fib = fibonacci(10);
	vctrprt(fib);

	return 0;
}
