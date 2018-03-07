#include <stdio.h>

void jaden_case(char *);

int main() {
	char str1[] = "3people unFollowed me for the last week";
	jaden_case(str1);	// "3people Unfollowed Me For The Last Week"
	printf("%s\n", str1);

	char str2[] = "viDeo kIl1ed radio 5taR";
	jaden_case(str2);	// "Video Kil1ed Radio 5tar"
	printf("%s", str2);

	return 0;
}

void jaden_case(char *str) {
	int i = 0, flag = 1;

	/* loop check using flag */
	/* if else statement */
	while(str[i] != '\0') {
		if(flag) {
			if('a' <= str[i] && str[i] <= 'z') {
				str[i] -= 32;
			}
			flag = 0;
		}
		else {
			if('A' <= str[i] && str[i] <= 'Z') {
				str[i] += 32;
			}
		}

		if(str[i] == ' ') { flag = 1; }
		i++;
	}
}
