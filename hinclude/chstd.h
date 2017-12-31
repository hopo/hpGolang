// ***********************
// *** For c api by hp ***
// ***********************


#ifndef CHSTD_H

int carrlen(char *arr) {
	int length = 0;
	while(1) {
		if(arr[length] == 0) break;
		length++;
	}
	return length;
}

#endif
