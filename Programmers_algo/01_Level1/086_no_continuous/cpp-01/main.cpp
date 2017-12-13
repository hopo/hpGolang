// ...ing

#include <iostream>
#include <vector>

#include "../../../../hpinclude/api_lab.h"

using namespace std;

vector<char> no_continous(char* cs) {
	vector<char> ret;
	for(int i = 0; ;i++) {
		if(cs[i] == 0) { break; } // terminator, cs end, '\0'
		else if(cs[i-1] != cs[i]) { ret.push_back(cs[i]); } // push_back no continous
	}
	return ret;
}

int main() {
	char ex1[] = "133303"; // [1 3 0 3]
	cout<< boxprt(no_continous(ex1)) << endl;

	char ex2[] = "47330";  // [4 7 3 0]
	cout << boxprt(no_continous(ex2)) << endl;
	return 0;
}
