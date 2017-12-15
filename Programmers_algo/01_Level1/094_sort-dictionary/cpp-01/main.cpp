#include <iostream>
#include <sstream>
#include <map>

using namespace std;

string sort_dictionary(map<string, int> ncard) {
	stringstream ss;
	ss << "(";
	for(auto& x : ncard) { // &  what???
		ss << "("  << x.first << ", " << x.second << ")" ;
	}
	ss << ")";

	return ss.str();
}

int main() {
	map<string, int> ncard = {{"김철수", 78}, {"이하나", 97}, {"정진원", 88}};
	cout << sort_dictionary(ncard) << endl;
}
