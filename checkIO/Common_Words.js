"use strict";

function commonWords(first, second) {
	var intersect = [];
	for(var f of first.split(",")){
		for(var s of second.split(",")){
			if(f == s){intersect.push(f)};
		};
	};
	return String(intersect.sort());
};

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(commonWords("hello,world", "hello,earth"), "hello", "Hello");
    assert.equal(commonWords("one,two,three", "four,five,six"), "", "Too different");
    assert.equal(commonWords("one,two,three", "four,five,one,two,six,three"), "one,three,two", "1 2 3");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
