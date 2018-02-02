#include <stdio.h>

int is_pair(char *);

int  main() {
	char *str1;
	str1 = "(hello)()";
	int ex1;
	ex1 = is_pair(str1);	/* 1, true */
	printf("%d\n", ex1);

	char *str2;
	str2 = ")(";
	int ex2;
	ex2 = is_pair(str2);	/* 0, false */
	printf("%d\n", ex2);

	char *str3;
	str3 = "(5+(40*(1+2)/2)+(1+2)*2)";
	int ex3;
	ex3 = is_pair(str3);	/* 1, true */
	printf("%d", ex3);
}

int is_pair(char *str) {
	int i = 0, b = 0;
	char box[32];

	// how many brackets
	while(str[i] != '\0') {
		if(str[i] == '(' || str[i] == ')') {
			box[b] = str[i];
			b++;

			if(b > 31) {	/* if box length over */
				printf("* TOO MANY BRACKETS * ");
				return -1;
			}
		}
		i++;
	}
	box[b] = '\0';
	
	// how many '('
	int chk = 0;
	i = 0;
	while(box[i] != '\0') {
		if(box[i] == '(') { chk++; }
		i++;
	}

	// 4 flag 
	int w, x, y, z;
	w = (b%2 == 0)? 1 : 0; 
	x = (box[0] == '(')? 1 : 0;
	y = (box[b-1] == ')')? 1 : 0;
	z = (b/2 == chk)? 1 : 0;

	int ret;
	ret = (w&x&y&z)? 1 : 0;
	
	return ret;
}
