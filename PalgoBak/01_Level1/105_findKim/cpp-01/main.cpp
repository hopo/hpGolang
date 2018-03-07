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
	int lnth = fnames.size();
	for(int i = 0; i < lnth; ++i) {
		if(fnames[i] == "Kim") { return i; }
	}
	return -1;
}
