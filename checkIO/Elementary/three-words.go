package main

import "regexp"

func threeWords(data string) {
	re := regexp.MustCompile("[a-zA-Z]+ [a-zA-Z]+ [a-zA-Z]+")
	mat := re.MatchString(data)
	println(mat)
}

func main() {
	threeWords("Hello World hello") //true
	threeWords("He is 123 man") //false
	threeWords("1 2 3 4") //false
    threeWords("bla bla bla bla") //true
	threeWords("Hi") //false
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
