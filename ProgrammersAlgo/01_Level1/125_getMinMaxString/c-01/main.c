#include <stdio.h>

typedef struct Data
{
	int min;
	int max;
} Data;

Data getMinMaxString(char str[]);

int main() {
	Data ex1 = getMinMaxString("1 2 3 4"); // 1, 4
	printf("%c, %c", ex1.min, ex1.max);

	return 0;
}

Data getMinMaxString(char str[]) {
	int min, max, i;
	min = max = str[0];
	for(i = 0; str[i] != '\0'; i++) {
		if('1' <= str[i] && str[i] <= '9') {
			if(min > str[i]) { min = str[i]; }
			if(max < str[i]) { max = str[i]; }
		}
	}

	Data dt;
	dt.min = min;
	dt.max = max;
	
	return dt;
}
