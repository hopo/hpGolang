package main

import (
    "fmt"
    "regexp"
)

func findMessage(data string) {
    re := regexp.MustCompile("[^A-Z]")
    fmt.Println(re.ReplaceAllString(data, ""))
}

func main() {
    findMessage("How are you? Eh, ok. Low or Lower? Ohhh.") //"HELLO", "hello"
    findMessage("hello world!") //"", "Nothing"
    findMessage("HELLO WORLD!!!")   //"HELLOWORLD", "Capitals"
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
