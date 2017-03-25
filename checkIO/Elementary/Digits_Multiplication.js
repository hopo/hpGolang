"use strict";

function digitsMultip(data) {
	var arr = Array.from(String(data));
	var z = 1;
	for(var i=0; i < arr.length; i++){
		if(arr[i] == 0){continue;};
		z = z*arr[i];
	};
    return z;
};

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(digitsMultip(123405), 120, "1st");
    assert.equal(digitsMultip(999), 729, "2nd");
    assert.equal(digitsMultip(1000), 1, "3rd");
    assert.equal(digitsMultip(1111), 1, "4th");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
