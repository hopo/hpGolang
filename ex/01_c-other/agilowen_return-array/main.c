// https://www.youtube.com/watch?v=EDy8BmAW3uo&t=1s

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SIZE 100
#define RETURN_STRING_LITERAL 0
#define RETURN_GLOABAL_ARR 1
#define RETURN_STATIC_ARR 0
#define RETURN_DYNAMIC_STRING 0
#define PREALLOCATED_BUFFER 0
#define RETURN_STRUCT_ARR 0

#include "array_returning_functions.c"

/********** MAIN **********/
int main(void) {
#if RETURN_STRING_LITERAL
	char *s;

	s = return_string_literal();
	printf("%s\n", s);

#elif RETURN_GLOABAL_ARR
	char *s;

	s = return_gloabal_arr();
	printf("%s\n", s);

#elif RETURN_STATIC_ARR
	char *s;

	s = return_static_arr();
	printf("%s\n", s);

#elif RETURN_DYNAMIC_STRING
	char *buffer;

	buffer = return_dynamic_string();
	printf("%s\n", buffer);
	if(buffer) { free(buffer); }

#elif PREALLOCATED_BUFFER 
	char *buffer;

	buffer = malloc(SIZE);
	if(buffer) {
		fill_preallocated_buffer(buffer, SIZE);
		printf("%s\n", buffer);
		free(buffer);
	}

#elif RETURN_STRUCT_ARR
	struct Data d;
	
	d = return_struct_arr();
	printf("%s\n", d.buffer);

#endif
	
	return 0;
}
