

/*
// ###Digits_Multiplication
function digitsMultip(data) {
	var arr = Array.from(String(data));
	var z = 1
	for(var i=0; i < arr.length; i++){
		if(arr[i] == 0){continue;};
		z = z*arr[i];
	};
    return z;
};

console.log(digitsMultip(3030));
*/

/*
// ###The_Most_Number : sort()
function mostNumbers(numbers){
	var arr = Array.from(arguments);
	return arr.sort(function(a, b){return a-b});
};

var nums = mostNumbers(6, 0, -3, 3);
console.log(nums);
*/


/*
// ###Tree_words : regular expression
function threeWords(data) {
	var words = data.split(" ");
	console.log(words);
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

var data = "ab ddb asadasb 88 a";
console.log(threeWords(data));
*/


/*
// ###asert.equal
var assert = require("assert");

function test(){
	assert.equal(1, 1, "3");
}

test();
*/


/*
// ###Secret_Message : case search
var str = "Good-Bye!";

function test(data){
	for(i=0; i < str.length; i++){
		if(data[i] != data[i].toLowerCase()){
			console.log(data[i]);
		};
	};
}

test(str);
*/
