#ifndef hp_h
#include <iostream>

class hp {
	public:
		hp(int a) {
			aaa = a;
		}
		void prt() {
			std::cout << aaa;	
		}
	private:
		int aaa;
};

#endif
