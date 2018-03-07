// ing
#include <stdio.h>

int jumpCase(int);

int main() {
	int n1 = 4, ex1;
	ex1 = jumpCase(n1); // 5
	printf("%d\n", ex1);

	/*
	int n2 = 7, ex2;
	ex2 = jumpCase(n2); // 21
	printf("%d", ex2);
	*/
}

int jumpCase(int n) {
	/* improve it */
	int i, ret, a = 0, b = 1;
	for(i = 0; i < n+1; i++) {
		ret = a + b;
		b = a;
		a = ret;
	}
	
	return ret;
}

/*
func jumpCase(n int) int {
	dd := dduim(n)                 // dduim make *ex: [2 2 1]
	sdd := [][]int{dd}             // data for two_breaker()
	vrdd_i := two_breaker(sdd)     // dd two breaker variety dduim slice
	vrdd_s := isl_conv_ssl(vrdd_i) // data conv for make_vari_sli()

	var ret int
	var x []string
	for i, _ := range vrdd_s {
		x = make_vari_sli(vrdd_s[i])
		if i != len(vrdd_s) {
			x = remove_samevalue(x)
		}
		ret += len(x)
	}
	return ret
}

func dduim(n int) []int {
	var isl []int
	for {
		if n == 2 {
			isl = append(isl, n)
			break
		}
		if n == 1 {
			isl = append(isl, n)
			break
		}
		if n-2 > 0 {
			isl = append(isl, 2)
			n -= 2
		}
	}
	return isl
}

func two_breaker(slsl [][]int) [][]int {
	ret := make([][]int, len(slsl))
	copy(ret, slsl)
	chk := make([]int, len(ret[len(ret)-1]))
	copy(chk, ret[len(ret)-1])
	var j int
	for i := 0; i < len(chk); i++ {
		j = len(chk) - 1 - i
		if chk[j] == 2 {
			chk[j] = 1
			chk = append(chk, 1)
			ret = append(ret, chk)
			return two_breaker(ret)
		}
	}
	return ret
}

func isl_conv_ssl(islsl [][]int) [][]string {
	var ret [][]string
	var ssl []string
	for _, dd := range islsl {
		for _, v := range dd {
			ssl = append(ssl, string(v+48))
		}
		ret = append(ret, ssl)
		ssl = []string{}
	}
	return ret
}

func make_vari_sli(ssl []string) []string {
	var fix string
	var ret, box []string
	box = ssl
	lnth := len(ssl)
	if lnth == 1 {
		return ssl
	}

	for i := 0; i < lnth; i++ {
		fix, box = box[0], box[1:] // set fix, box

		var send []string
		send = make([]string, len(box))
		copy(send, box)
		vri := make_vari_sli(send)

		var rcv string
		for _, v := range vri {
			rcv = fix + v
			ret = append(ret, rcv)
			rcv = ""
		}

		box = append(box, fix) // init: fix, box
	}

	return ret
}

func remove_samevalue(ssl []string) []string {
	var box, ret []string
	box = ssl
	lnth := len(box)
	var tag string
	for i := 0; i < lnth; i++ {
		tag = box[i] // set
		for j := i; j < lnth; j++ {
			if i != j && tag == box[j] {
				box[i] = ""
				break
			}
		}
	}
	for _, v := range box {
		if v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}
*/
