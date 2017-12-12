#include <iostream>
#include <vector>

#include "../../../../hpinclude/api_lab.h"

using namespace std;

vector<int> divisible(vector<int> smpl, int d) {
	vector<int> ret;
	for(auto v : smpl) {
		if(v % d == 0) { ret.push_back(v); }
	}
	return ret;
}

int main() {
	vector<int> ex1{5, 9, 7, 10}; int n = 5; // [5 10]
	cout << boxprt(divisible(ex1, n)) << endl;
}
