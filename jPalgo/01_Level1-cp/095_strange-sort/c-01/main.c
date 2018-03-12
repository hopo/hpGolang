#include <stdio.h>

#define ALP "abcdefghijklmnopqrstuvwxyz"

typedef char *String;

String *strange_sort(String words[], int n);

int main() {
	String words[] = {"sun", "bed", "car"}; // car, bed, sun
	String *ex1 = strange_sort(words, 1);
	printf("%s, %s, %s", ex1[0], ex1[1], ex1[2]);

	return 0;
}

String *strange_sort(String words[], int n) {
	int i, r = 0, j = 0;
	String box;
	String ret[3];
	while(ALP[j] != '\0') {
		for(i = 0; i < 3; i++) {
			box = *(words+i);
			if(ALP[j] == box[n]) { ret[r] = box; r++; };
		}
		j++;
	}

	String *dt = ret;

	return dt;
}
