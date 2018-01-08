// ing
// char array append
// receive pointer make char array

#include <stdio.h>

char* reverseStr(char*);

int main() {
	char* ex1 = reverseStr("Zbcdefg");     // "gfedcbZ"
	printf("%c", ex1[2]);

	return 0;
}

char* reverseStr(char* str) {
	int lnth = 0;

	while(*(str+lnth) != '\0') { lnth++; }
	char rev[lnth];
	
	int i, j;
	for(i = 0; i < lnth; i++) {
		j = lnth-1-i;
		rev[i] = *(str+j);
	}

	char* dt;
	dt = rev;

	return dt;
}




/*

func main() {
	ex1 := reverseStr("Zbcdefg")     // "gfedcbZ"
	ex2 := reverseStr("deSadFklJfg") // "lkgfeddaSJF"
}

func reverseStr(s string) (ret string) {
	for i := 0; i < 200; i++ {
		for _, v := range s {
			if i == int(v) {
				ret = string(v) + ret
			}
		}
	}
	return
}
*/
