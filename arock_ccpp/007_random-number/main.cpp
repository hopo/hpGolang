// https://www.youtube.com/watch?v=xJtmRDDpQIA&list=PL4SIC1d_ab-b4zy_3FDRIiohszShOZ0PK&index=7

#include <iostream>
#include <ctime>

using namespace std;

int main() {
	// random-number
	srand((unsigned int)time(0));	

	cout << rand() << endl;
	cout << rand() << endl;
	cout << rand() << endl;

	cout << endl;

	cout << (rand() % 100) << endl; // 0~99
	cout << (rand() % 101 + 100) << endl; // 100~200
	cout << (rand() % 10000 / 100.f) << endl; // xx.xx

	cout << endl;

	int iUpgrade = 0;
	cout << "Upgrade number???: ";
	cin >> iUpgrade;

	// up chance;
	float fPercent = rand() % 10000 / 100.f;

	// 0~3 : 100%;  4~6: 40%; 7~9: 10%;
	// 10~13: 1%; 14~15: 0.01%;
	cout << "Upgrade: " << iUpgrade << endl;
	cout << "Percent: " << fPercent << endl;
	
	if(iUpgrade <= 3)
		cout << "UP COMPLETE" << endl;
	else if(4 <= iUpgrade && iUpgrade <= 6) {
		if(fPercent < 40.f)
			cout << "UP COMPLETE" << endl;
		else 
			cout << "FAILED" << endl;
	}
	else if(7 <= iUpgrade && iUpgrade <= 9) {
		if(fPercent < 10.f)
			cout << "UP COMPLETE" << endl;
		else 
			cout << "FAILED" << endl;
	}
	else if(10 <= iUpgrade && iUpgrade <= 13) {
		if(fPercent < 1.f)
			cout << "UP COMPLETE" << endl;
		else 
			cout << "FAILED" << endl;
	}
	else if(14 <= iUpgrade && iUpgrade <= 15) {
		if(fPercent < 0.01f)
			cout << "UP COMPLETE" << endl;
		else 
			cout << "FAILED" << endl;
	}

	return 0;
}
