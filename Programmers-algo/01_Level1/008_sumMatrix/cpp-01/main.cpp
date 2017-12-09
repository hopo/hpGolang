// https://programmers.co.kr/learn/challenge_codes/148

#include <iostream>
#include <vector>

using namespace std;

vector<vector<int>> sumMatrix(vector<vector<int>>A, vector<vector<int>>B) {
	vector<vector<int>> answer;

	return answer;
}

int main() {
	vector<vector<int>> a{{1,2},{2,3}}, b{{3,4},{5,6}}; // {{4, 6}, {{7, 9}}
	vector<vector<int>> answer = sumMatrix(a,b);

	for(int i = 0; i < answer.size(); i++) {
		for(int j = 0; j < answer[0].size(); j++) {
			cout << answer[i][j] << " ";
		}
		cout << "\n";
	}
}
