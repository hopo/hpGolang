// ing
#include <iostream>
#include <vector>

using namespace std;

int findKim(vector<string>);

int main() {
	vector<string> ex1{"Queen", "Tod", "Kim"};	// 2
	cout << findKim(ex1) << endl;

	vector<string> ex2{"Gim", "Obama", "xi"};	// -1
	cout << findKim(ex2) << endl;

	return 0;
}

int findKim(vector<string> fnames) {
	int ret = -1;
	/*
	for i, v := range fnames {
		if v == "Kim" {
			ret = i
		}
	}
	*/
	return ret;
}
