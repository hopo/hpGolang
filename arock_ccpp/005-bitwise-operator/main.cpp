#include <iostream>

using namespace std;

int main() {
	const int iAttack = 0x00000001;	//	00001
	const int iArmor = 0x00000002;	// 00010
	const int iHP = 0x00000004;	// 00100
	const int iMP = 0x00000008;	// 01000
	const int iCritical = 0x00000010;	// 10000
	
	// OR oper: 00001 | 00100 = 00101 | 10000 = 10101
	int iBuf = iAttack | iHP | iCritical;	// 10101

	// XOR oper: 10101 ^ 00100 = 10001
	iBuf ^= iHP; // on/off
	// XOR oper: 10001 ^ 00100 = 10101
	//iBuf ^= iHP;

	// AND oper:  10101 & 00001 = 00001
	cout << "Attack: " <<  (iBuf & iAttack) << endl;	// true, 1
	// 10101 & 00010 = 00000
	cout << "Armor: " <<  (iBuf & iArmor) << endl;	// false
	cout << "HP: " <<  (iBuf & iHP) << endl;	// false
	cout << "MP: " <<  (iBuf & iMP) << endl;	// false
	cout << "Critical: " <<  (iBuf & iCritical) << endl;	// true, 16

	cout << endl;

	/*
	 * shift oper: >>, <<
	 * 20 << 2 = 80  ; 20*(2**2)
	 * 	10100 -> 1010000
	 * 20 << 3 = 160 ; 20*(2**3)
	 * 	10100 -> 10100000
	 * 20 << 4 = 320 ; 20*(2**4)
	 * 	10100 -> 101000000
	 *
	 * 20 >> 2 = 5  ; 20/(2**2)
	 *  10100 -> 101	
	 * 20 >> 3 = 2  ; 20/(2**3)
	 *  10100 -> 10	
	 * 20 >> 4 = 1  ; 20/(2**4)
	 *  10100 -> 1	
	 */

	int iHigh = 187;
	int iLow = 13560;

	int iNumber = iHigh;
	// (187) ; type int is 32bit(4bytes)

	// 187 in iNumber.
	// this value shift 16bit
	iNumber <<= 16;
	// (187) - 0000000000000000
	
	// fill in low 16bit
	iNumber |= iLow;
	// (187) - (13560)
	
	// print High and Low
	cout << "High: " << (iNumber >> 16) << endl;
	cout << "Low: " << (iNumber & 0x0000ffff) << endl;

	cout << endl;

	// prefix and postfix
	iNumber = 10;

	// prefix
	++iNumber;

	// postfix
	iNumber++;

	cout << ++iNumber << endl; // after plus, print
	cout << iNumber++ << endl; // after print, plus
	cout << iNumber << endl;

	return 0;
}
