// ...ing

#include <stdio.h>
#include <string.h>

typedef struct Namecard
{
	char name[20];
	int number;
} Ncard;

Ncard sort_dictionary(Ncard);

int main() {
	Ncard n1, n2, n3;
	strcpy(n1.name, "김철수"); n1.number = 78;  
	strcpy(n2.name, "이하나"); n2.number = 97;
	strcpy(n3.name, "정진원"); n3.number = 88;
	
	Ncard cards[] = {n1, n2, n3};

	for(int i = 0; i < 3; i++) {
		printf("%s, ", cards[i].name);
		printf("%d\n", cards[i].number);
	}

	return 0;
}


/*
func sort_dictionary(nc map[string]int) (ret []Ncard) {
	var slis []string
	for k, _ := range nc {
		slis = append(slis, k)
	}
	sort.Strings(slis)
	for _, k := range slis {
		ret = append(ret, Ncard{k, nc[k]})
	}
	return
}
*/
