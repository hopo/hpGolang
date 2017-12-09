#include <iostream>
#include <vector>

using namespace std;

int main() {
	vector<vector<int>>::iterator mmm;

	int m;
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			m = i * j;
			mmm[i][j] = m;
			cout << "i" << i << ": " << mmm[i][j] << ", ";
		}
	}

	return 0;
}
