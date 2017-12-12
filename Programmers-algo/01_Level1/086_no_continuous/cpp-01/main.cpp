// ...ing

#include <iostream>
#include <vector>

#include "../../../api_lab.h"

using namespace std;

void no_continous(char* cs) {
	cout << cs << endl;

	/*
	for (i, v : range s) {
		n, e := strconv.Atoi(string(v))
		if (e != nil) { panic(e) }
		if (i != 0 && s[i-1] == s[i]) { continue }
		else { ret = ret.push_back(n); }
	}
	*/
}

int main() {
	char ex1[] = "133303"; // [1 3 0 3]
	no_continous(ex1);
	//char ex2[] = "47330"  // [4 7 3 0]
	return 0;
}
