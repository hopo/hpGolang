package main

import (
    "fmt"
    "regexp"
)

func findMessage(data string) string {
    return regexp.MustCompile("[^A-Z]").ReplaceAllString(data, "")
}

func main() {
    fmt.Println(findMessage("How are you? Eh, ok. Low or Lower? Ohhh.")) //"HELLO", "hello"
    fmt.Println(findMessage("hello world!")) //"", "Nothing"
    fmt.Println(findMessage("HELLO WORLD!!!"))   //"HELLOWORLD", "Capitals"
}

/*
function findMessage(data) {
    var pars = '';
	for(var i = 0; i < data.length; i++){
		if(data[i] != data[i].toLowerCase()){
			pars += data[i];
		};
	};
	if(pars == "HELLO"){
		return pars;
	}else if(!pars){
		return pars;
	}
	return pars;
};
*/

/*
var findMessage = data => data.replace(/[^A-Z]/g, '');
*/
