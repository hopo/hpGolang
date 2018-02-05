#include <stdio.h>

/*
// -std=c11
#define type_name(T)	\
	_Generic((T),\
		char: "char",\
		int: "int", \
		float: "float"\
		default: "other" )
*/

	
int main() {
	int i = 8;
	float f = 4.0;

	printf("int / float\n");

	printf("ret_d: %f\n", i/f);
	printf("ret_p: %f\n", i/f);
	printf("ret_f: %f\n", i/f);
	printf("ret_x: %f\n", i/f);

	return 0;
}
