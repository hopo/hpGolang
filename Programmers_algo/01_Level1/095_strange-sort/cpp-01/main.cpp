#include <iostream>
#include <vector>
#include "../../../../hinclude/hlib.h"

using namespace std;

const string sALP = "abcdefghijklmnopqrstuvwxyz";

vector<string> strange_sort(vector<string>, int);

int main() {
	vector<string> ex1 = {"sun", "bed", "car"}; int n1 = 1;	// [car bed sun]
	cout << boxprt(strange_sort(ex1, n1)) << endl;

	return 0;
}


vector<string> strange_sort(vector<string> ex1, int n1) {
	vector<string> ret;
	for(auto& v : sALP) {
		for(auto& w : ex1) {
			if(v == w[1]) {
				ret.push_back(w);
			}
		}
	}

	return ret;
}
