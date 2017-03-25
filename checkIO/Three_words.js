// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions

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

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(threeWords("Hello World hello"), true, "1st example");
    assert.equal(threeWords("He is 123 man"), false, "2nd example");
    assert.equal(threeWords("1 2 3 4"), false, "3rd example");
    assert.equal(threeWords("bla bla bla bla"), true, "4th example");
    assert.equal(threeWords("Hi"), false, "Letters");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
