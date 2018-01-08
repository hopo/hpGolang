#include <stdio.h>

#define tokenpaster(n) printf("token" #n " = %d", token##n)

main(void) {
	int token34 = 40;
	tokenpaster(34);
	
	// confer
	printf("\n*** token34 = %d", token34);

	return 0;
}
