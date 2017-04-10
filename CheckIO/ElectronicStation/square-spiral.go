"use strict";

function findDistance(first, second) {
    return 0;
}

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(findDistance(1, 9), 2, "1st example");
    assert.equal(findDistance(9, 1), 2, "2nd example");
    assert.equal(findDistance(10, 25), 1, "3rd example");
    assert.equal(findDistance(5, 9), 4, "4th example");
    assert.equal(findDistance(26, 31), 5, "5th example");
    assert.equal(findDistance(50, 16), 10, "6th example");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
