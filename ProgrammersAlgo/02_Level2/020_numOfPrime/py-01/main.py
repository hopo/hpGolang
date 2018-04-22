# ...ing

def num_of_prime(num):
    
    return -1

if __name__ == '__main__':
    ex1 = num_of_prime(10) # 4, [2 3 5 7]
    print(ex1)

    ex2 = num_of_prime(5) # 3, [2 3 5]
    print(ex2)

"""
func numOfPrime(n int) int {
	r := makenop(n)
	return len(r)
}

func makenop(lmt int) []int {
	var isl []int
	for i := 2; i < lmt+1; i++ {
		isl = func(isl []int, i int) []int {
			for _, v := range isl {
				if i%v == 0 {
					return isl
				}
			}
			return append(isl, i)
		}(isl, i)
	}
	return isl
}
"""
