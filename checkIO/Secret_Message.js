"use strict";

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

/*
var findMessage = data => data.replace(/[^A-Z]/g, '');
*/

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(findMessage("How are you? Eh, ok. Low or Lower? Ohhh."), "HELLO", "hello");
    assert.equal(findMessage("hello world!"), "", "Nothing");
    assert.equal(findMessage("HELLO WORLD!!!"), "HELLOWORLD", "Capitals");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
