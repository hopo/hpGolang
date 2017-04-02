package main

import (
    "fmt"
    "regexp"
)

func threeWords(data string) bool {
	return regexp.MustCompile("[a-zA-Z]+ [a-zA-Z]+ [a-zA-Z]+").MatchString(data)
}

func main() {
	fmt.Println(threeWords("Hello World hello")) //true
	fmt.Println(threeWords("He is 123 man")) //false
	fmt.Println(threeWords("1 2 3 4")) //false
    fmt.Println(threeWords("bla bla bla bla")) //true
	fmt.Println(threeWords("Hi")) //false
}

/*
function threeWords(data) {
    var words = data.split(" ");
	var rslt = "";
	for(i = 0; i < words.length; i++){
		if(parseInt(words[i])){
			rslt += "n";
		}else{
			rslt +="w";
		};
	};
	if(rslt.match(/www/)){
		return true;
	}else{
		return false;
	};
};

/*
function threeWords(data) {
    return /[a-z]+ [a-z]+ [a-z]+/ig.test(data);
}
*/
