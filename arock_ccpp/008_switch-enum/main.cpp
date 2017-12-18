// https://www.youtube.com/watch?v=vLxGf_QNv5s&index=8&list=PL4SIC1d_ab-b4zy_3FDRIiohszShOZ0PK

#include <iostream>

using namespace std;

// const
enum NUM
{
	NUM_0, // no commnet, ~++	
	NUM_1,	
	NUM_2,	
	NUM_3,	
};

// const
#define NUM_4 4

int main() {
	int iNumber;
	cout << "Enter number: ";
	cin >> iNumber;

	switch(iNumber) {
		case NUM_1:
			cout << "ONE!" << endl;
			break;
		case NUM_2:
			cout << "TWO!" << endl;
			break;
		case NUM_3:
			cout << "THREE!" << endl;
			break;
		case NUM_4:
			cout << "FOUR!" << endl;
			break;
		default:
			cout << "etc..." << endl;
			break;
	}

	// type enum declare. that is 4bytes
	NUM eNum;  

	cout << "sizeof(NUM): " << sizeof(NUM) << endl;
	cout << "typeid(eNum).name(): " << typeid(eNum).name() << endl;

	return 0;
}
