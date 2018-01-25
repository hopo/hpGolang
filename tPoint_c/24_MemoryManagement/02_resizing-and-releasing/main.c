#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
	char name[100];
	char *description;

	strcpy(name, "Zara Ali");

	/* alocate memory dynamically */
	description = malloc(30 * sizeof(char));

	if(description == NULL) {
		fprintf(stderr, "Error - unable to allocate required memory\n");
	}
	else {
		strcat(description, "She is in calss 10th");
	}

	printf("Name = %s\n", name);
	printf("Description: %s\n",description);

	/* release memory using free() function */
	free(description);
}
