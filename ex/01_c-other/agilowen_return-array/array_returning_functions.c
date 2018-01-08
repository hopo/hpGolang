/* (1) String Literal */
char *return_string_literal() {
	char *string_literal = "(1) This is a string literal";

	return string_literal;
}

/* (2) Global Array */
char global_arr[SIZE];
char *return_gloabal_arr() {
	strncpy(global_arr, "(2) This string is in a global array", SIZE);

	return global_arr;
}

/* (3) Static Array */
char *return_static_arr() {
	static char static_arr[SIZE];

	strncpy(static_arr, "(3) This string is in a static array", SIZE);

	return static_arr;
}

/* (4) Explictly allocate memory */
char *return_dynamic_string() {
	char *buffer = malloc(SIZE);

	strncpy(buffer, "(4) This string is in a malloc'ed buffer", SIZE);

	return buffer;
}

/* (5) Caller allocates/frees memory */
void fill_preallovated_buffer(char *result, int size) {
	strncpy(result, "(5) This string is in a preallocated buffer", size);
}

/* (6) Wrap fixed-size array in a struct */
struct Data {
	char buffer[SIZE];
};
struct Data return_struct_arr() {
	struct Data d;

	strncpy(d.buffer, "(6) This string is in a struct-wrapped array", SIZE);

	return d;
}
