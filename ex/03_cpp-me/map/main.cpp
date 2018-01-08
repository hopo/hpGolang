#include <iostream>
#include <string>
#include <map>

using namespace std;

int main() {
	map<string, int> ncard;

	ncard["하루바"] = 27;
	ncard["남숙희"] = 18;
	ncard["Obama"] = 98;
	ncard["Amber"] = 56;

	map<string, int>::iterator it;

	for(it = ncard.begin(); it != ncard.end(); ++it) {
		cout <</* it->first <<*/ ": " << it->second << "\n";
	}
	return 0;
}
