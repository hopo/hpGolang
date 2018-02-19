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
	while(str[i] != '\0') {
		switch(flag) {
			case 0:
				if('A' <= str[i] && str[i] <= 'Z') { str[i] += 32; }
				break;
			case 1:
				if('a' <= str[i] && str[i] <= 'z') { str[i] -= 32; }
				flag = 0;
		}

		if(str[i] == ' ') { flag = 1; }
		i++;
	}
}
