#include <iostream>

using namespace std;

class Box {
	public:
		double length; // Length of box
		double breadth; // Breadth of box
		double height; // Height of box

		// Member functions declaraion
		double getVolume(void);
		void setLength(double len);
		void setBreadth(double bre);
		void setHeight(double hei);
};

// Member functions definitions
double Box::getVolume(void) {
	return length * breadth * height;
}

void Box::setLength(double len) {
	length =len;
}

void Box::setBreadth(double bre) {
	breadth = bre;
}

void Box::setHeight(double hei) {
	height = hei;
}

// Main functions for the program
int main() {
	Box Box1; // Declare Box1 of type Box
	Box Box2; // Declare Box2 of type Box
	double volume = 0.0;

	// box 1 sepcification
	Box1.setLength(6.0);
	Box1.setBreadth(7.0);
	Box1.setHeight(5.0);

	// box 2 sepcification
	Box2.setLength(12.0);
	Box2.setBreadth(13.0);
	Box2.setHeight(10.0);

	// volume of box 1
	volume = Box1.getVolume();
	cout << "Volume of Boc1 : " << volume << endl;

	// volume of box 2
	volume = Box2.getVolume();
	cout << "Volume of Box2 : " << volume << endl;
	return 0;
}
