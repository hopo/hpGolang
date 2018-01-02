// ..ing

#include <stdio.h>
#include <string.h>

#define ALP "abcdefghijklmnopqrstuvwxyz"

char* strange_sort(char* words[], int n) {
	int i, r = 0, j = 0;
	char* box;
	char* ret[3];
	while(ALP[j] != '\0') {
		for(i = 0; i < 3; i++) {
			box = *(words+i);
			if(ALP[j] == box[n]) { ret[r] = box; r++; };
		}
		j++;
	}
	
	printf("func() %s %s %s\n", ret[0], ret[1], ret[2]);
	
	char* dt = ret[3];

	return dt;
}

int main() {
	char* words[] = {"sun", "bed", "car"}; // car, bed, sun
	char* ex = strange_sort(words, 1);

	printf("main() %s\n", &(ex));

	return 0;
}

