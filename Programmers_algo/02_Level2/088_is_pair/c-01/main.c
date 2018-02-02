#include <stdio.h>

int is_pair(char *);

int  main() {
	char *str1;
	str1 = "(hello)()";	// 1, true
	int ex1;
	ex1 = is_pair(str1);

	/*
	ex2 := is_pair(")(")                       // false
	ex3 := is_pair("(5+(40*(1+2)/2)+(1+2)*2)") // true
	fmt.Println(ex1)
	fmt.Println(ex2)
	fmt.Println(ex3)
	*/
}

int is_pair(char *str) {
	printf("%s", str);
	
	return -1;

	/*
	bsl := []byte(s)
	var smpl []byte
	for _, v := range bsl {
		if v == 40 || v == 41 {
			smpl = append(smpl, v)
		}
	}
	l := len(smpl)
	var chk int
	for _, v := range smpl {
		if v == 40 {
			chk++
		}
	}
	a := l%2 == 0
	b := smpl[0] == 40
	c := smpl[l-1] == 41
	d := chk == l/2
	if a && b && c && d {
		return true
	}

	return false
	*/
}
