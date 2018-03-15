#include <stdio.h>
#include <string.h>

const int mdays[12] = {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
const char *dnames[] = {"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"};

void getDayName(int, int, char *);

/* default: January 1st is Friday. */

int main() {
	int m1 = 5, d1 = 24;
	char ret1[8];
	getDayName(m1, d1, ret1); // TUE
	printf("%s\n", ret1);
	
	int m2 = 11, d2 = 5;
	char ret2[8];
	getDayName(m2, d2, ret2); // SAT
	printf("%s", ret2);
}

void getDayName(int month, int day, char *ret) {
	int i, r;
	/* how many total days */
	for(i = 0; i < month-1; i++) { day += mdays[i]; }

	r = day%7;

	/* what is day names */
	if(r < 3) { strcpy(ret, dnames[r+4]); }
	else { strcpy(ret, dnames[r-3]); }
}
