#include <stdio.h>

typedef struct Namecard
{
	char[20] name,
	int number
};

Namecard sort_dictionary(Namecard);

int main() {
	Namecard ncard[] = {
		{"김철수", 78},
		{"이하나", 97},
		{"정진원", 88}
	};
	Namecard ex1 = sort_dictionary(ncard);

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
