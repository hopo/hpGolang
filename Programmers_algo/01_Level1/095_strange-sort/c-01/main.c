// ..ing

#include <stdio.h>
#include <string.h>

#define ALP "abcdefghijklmnopqrstuvwxyz"

typedef struct String
{
	char word[10];	
} str;

int main() {
	str a, b, c;
	strcpy(a.word, "sun");
	strcpy(b.word, "bed");
	strcpy(c.word, "car");

	printf("a.word: %s", a.word);
	printf("\nALP[6]: %c", ALP[6]);

	//ex := strange_sort(words, 1) // [car bed sun]

	return 0;
}

/*
func strange_sort(ssl []string, n int) (ret []string) {
	for _, r := range ALP {
		for _, v := range ssl {
			if byte(r) == v[n] {
				ret = append(ret, v)
			}
		}
	}
	return
}
*/
