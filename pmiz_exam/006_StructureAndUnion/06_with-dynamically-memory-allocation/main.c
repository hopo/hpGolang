#include <stdio.h>
#include <stdlib.h>

struct course
{
	int marks;
	char subject[30];
};

int main() {
	struct course *ptr;
	int i, noOfRecords;

	printf("Enter number of recods: ");
	scanf("%d", &noOfRecords);

	// Allocates the memory for noOfRecodes structures with pointer ptr pointing to the base address.
	ptr = (struct course*) malloc(noOfRecords*sizeof(struct course));

	for(i = 0; i < noOfRecords; ++i) {
		printf("Enter name fo the subject and marks respectively:\n");
		scanf("%s %d", &(ptr+i)->subject, &(ptr+i)->marks);
	}

	printf("Displaying Information:\n");

	for(i = 0; i < noOfRecords; ++i) {
		printf("%s\t%d\n", (ptr+i)->subject, (ptr+i)->marks);
	}

	return 0;
}
