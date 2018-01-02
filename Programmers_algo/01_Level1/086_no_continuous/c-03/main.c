#include <stdio.h>
#include <string.h>
//#include "/Users/HPMBA/workspace/hinclude/chstd.h"
#include "/home/hp/workspace/hinclude/chstd.h"

int* no_continous(char*);

int main() {
	char data1[] = "133303";	// 1 3 0 3
	int* ex1 = no_continous(data1);
	printf("%d %d %d %d", *ex1, *(ex1+1), *(ex1+2), *(ex1+3));

	printf("\n");

	char data2[] = "47330";  // 4 7 3 0
	int* ex2 = no_continous(data2);
	printf("%d %d %d %d", *ex2, *(ex2+1), *(ex2+2), *(ex2+3));

	return 0;
}

int* no_continous(char* data) {
	int len, i, j = 1;
	len = carrlen(data);

	char box[len];
	box[0] = data[0];
	
	for(i = 1; i < len; i++) {
		if(data[i-1] != data[i]) {
			box[j] = data[i];
			j++;
		};
	}

	int ret[j];	
	for(i = 0; i < j; i++) {
		ret[i] = box[i]-'0';
	}

	int* dt = ret;

	return dt;
}

/*
func no_continous(s string) (ret []int) {
	for i, v := range s {
		n, e := strconv.Atoi(string(v))
		if e != nil {
			panic(e)
		}
		if i != 0 && s[i-1] == s[i] {
			continue
		} else {
			ret = append(ret, n)
		}
	}
	return
}
*/
