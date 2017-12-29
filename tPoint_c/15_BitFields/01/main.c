#include <stdio.h>
#include <string.h>

/* define simple structure */
struct
{
	unsigned int widthValidated;
	unsigned int heightValidated;
} status1;

/* define a structure with bit fields */
struct {
	unsigned int widthValidated : 1;
	unsigned int heightValidated : 1;
} status2;

int main() {
	printf("Memory size ocuppied by status1 : %d\n", sizeof(status1));
	printf("Memory size ocuppied by status2 : %d\n", sizeof(status2));

	return 0;
}
