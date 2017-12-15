#include <iostream>
#include <vector>

using namespace std;

const string sALP = "abcdefghijklmnopqrstuvwxyz";

int strange_sort(vector<string>, int);

int main() {
	vector<string> ex1 = {"sun", "bed", "car"}; int n1 = 1;	// [car bed sun]
	cout << strange_sort(ex1, n1) << endl;

	return 0;
}


int strange_sort(vector<string> ex1, int n1) {
	for(auto& v : sALP ) {
		cout << v;
	}

	/*
	for _, r := range ALP {
		for _, v := range ssl {
			if byte(r) == v[n] {
				ret = append(ret, v)
			}
		}
	}
	return
	*/

	cout << endl;
	return -1;
}
