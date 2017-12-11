#include <iostream>
#include <vector>
#include <sstream>

using namespace std;

template <class T>
string boxprt(T box) {
	stringstream ss; // sstring::stringstream
	ss << "[";
	for (int i = 0; i < box.size(); i++) { 
		if (i == box.size()-1) {
			ss << box[i];
			break;
		}
		ss << box[i] << " ";
	}
	ss << "]";
	return ss.str(); // sstring::stringstream::str()
}

vector<char> string_middle(char* cs) {
	int csize;
	for (int i = 0; ; i++) {
		if (cs[i] == 0) {
			csize = i;
			break;
		}
	}
	
	vector<char> ret;
	if (csize % 2 == 0) { ret.push_back(cs[csize/2-1]); }
	ret.push_back(cs[csize/2]);

	return ret;
}

int main() {
	char ex1[] = "power"; // [w]
	cout << boxprt(string_middle(ex1)) << endl;

	char ex2[] = "test";  // [e s]
	cout << boxprt(string_middle(ex2)) << endl;

	return 0;
}
