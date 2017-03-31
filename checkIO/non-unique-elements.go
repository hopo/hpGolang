"use strict";

function nonUniqueElements(data) {
	for(var i=0; i<data.length; i++) {
		var sock = data[i];
		delete data[i];
		data.includes(sock) ? data[i] = sock : null;
	};
	return data.filter(() => {return true});
}

/*
nonUniqueElements = a => a.filter(e => a.indexOf(e)!= a.lastIndexOf(e));
*/

var assert = require('assert');

if (!global.is_checking) {
    assert.deepEqual(nonUniqueElements([1, 2, 3, 1, 3]), [1, 3, 1, 3], "1st example");
    assert.deepEqual(nonUniqueElements([1, 2, 3, 4, 5]), [], "2nd example");
    assert.deepEqual(nonUniqueElements([5, 5, 5, 5, 5]), [5, 5, 5, 5, 5], "3rd example");
    assert.deepEqual(nonUniqueElements([10, 9, 10, 10, 9, 8]), [10, 9, 10, 10, 9], "4th example");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}