// *********
// hpkg pygo
// *********

package hpkg

import (
//"strconv"
)

// range(a, b): a to b
func Range(a, b int) (rslt []int) {
	for ; a < b; a++ {
		rslt = append(rslt, a)
	}
	return
}

// sum Float slice
func Sum_fsl(fsl []float64) (ret float64) {
	for _, v := range fsl {
		ret += v
	}
	return
}

// sum Int slice
func Sum_isl(isl []int) (ret int) {
	for _, v := range isl {
		ret += v
	}
	return
}

// value in list(slice)
func Value_in_isl(v int, is []int) (rslt []int) {
	for i, val := range is {
		if val == v {
			rslt = append(rslt, i)
		}
	}
	return
}

// check lower case by byte
func IsLower(b byte) bool {
	// byte 97~122 is lower case
	for i := 97; i < 123; i++ {
		if int(b) == i {
			return true
		}
	}
	return false
}

// check upper case by byte
func IsUpper(b byte) bool {
	// byte 65~90 is upper case
	for i := 65; i < 91; i++ {
		if int(b) == i {
			return true
		}
	}
	return false
}

// check digit case by byte
func IsDigit(b byte) bool {
	// byte 48~57 is digit
	for i := 48; i < 58; i++ {
		if int(b) == i {
			return true
		}
	}
	return false
}
