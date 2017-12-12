#include <iostream>
#include <vector>

using namespace std;

int main() {
	vector<vector<int>>::iterator vtm;

	int m;
	for (int i = 0; i < 2; i++) {
		for (int j = 0; j < 2; j++) {
			m = i * j;
			vtm[i][j] = m;
			cout << "i" << i << ": " << vtm[i][j] << "\n";
		}
	}

	return 0;
}
