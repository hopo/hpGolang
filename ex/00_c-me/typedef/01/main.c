#include <stdio.h>

typedef char* String;

int main() {
	String str;

	str = "good bye";

	printf("String str = \"%s\";", str);
	
	String w1, w2, w3;

	w1 = "my";
	w2 = "dear";
	w3 = "!";

	String words[] = {str, w1, w2, w3};

	printf("\n%s", words[0]);
	printf(" %s", words[1]);
	printf(" %s", words[2]);
	printf("%s", words[3]);

	printf("\nsizeof(words): %lu", sizeof(words));
	printf("\nsizeof(words[1]): %lu", sizeof(words[1]));
	printf("\nsizeof((words[1])[0]): %lu", sizeof((words[1])[0]));


	return 0;
}
