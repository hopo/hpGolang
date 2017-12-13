// https://programmers.co.kr/learn/challenge_codes/148

#include <iostream>
#include <vector>

using namespace std;

vector<vector<int>> sumMatrix(vector<vector<int>>A, vector<vector<int>>B) {
	vector<vector<int>> answer = {{0, 0}, {0, 0}};
	size_t st = A.size();

	for (int i = 0; i < st; i++) {
		for (int j = 0; j < st; j++) {
			answer[i][j] = A[i][j]+B[i][j];
		}
	}	
	return answer;
}

int main() {
	vector<vector<int>> a{{1, 2}, {2, 3}}, b{{3, 4}, {5, 6}}; // {{4, 6}, {{7, 9}}
	vector<vector<int>> answer = sumMatrix(a,b);

	for(int i = 0; i < answer.size(); i++) {
		for(int j = 0; j < answer[0].size(); j++) {
			cout << answer[i][j] << " ";
		}
		cout << "\n";
	}
}
