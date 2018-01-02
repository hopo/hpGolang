// ***********************
// *** For C api by hp ***
// ***********************


#ifndef CHSTD_H

// methods
int carrlen(char*);


// carrlen()
// p: char*
// r: int
int carrlen(char* arr) {
	int length = 0;
	while(arr[length] != '\0') { length++; }
	return length;
}

#endif
